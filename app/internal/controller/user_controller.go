package controller

import (
	"github.com/Vityalimbaev/Example-Backend/config"
	"github.com/Vityalimbaev/Example-Backend/internal/constants"
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/service/user"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type UserController struct {
	userService user.ServiceI
}

func NewUserController(userService user.ServiceI) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (controller *UserController) RegisterHandlers(private, public *gin.RouterGroup) {
	public.POST("/login", controller.LoginHandler)
	public.POST("/save-user", controller.SaveUserHandler)
	private.GET("/isauth", controller.IsAuthHandler)
	public.GET("/refresh-token", controller.RefreshTokenHandler)
	private.POST("/get-users", controller.GetUsersHandler)
}

// IsAuthHandler godoc
// @Summary Checking the access token for validity
// @Description  Get request with access token in header return success or error
// @Security BearerAuth
// @Tags         User
// @Success      200
// @Router       /isauth [get]
func (controller *UserController) IsAuthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, entity.ResponseView{
		IsSuccess: true,
		Data:      struct{}{},
	})
}

// LoginHandler godoc
// @Summary      SignIn by email and password
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 User body  entity.LoginForm true  "only email"
// @Success      200  {object}  entity.ResponseView{data=entity.TokensResponseView}
// @Router       /login [post]
func (controller *UserController) LoginHandler(ctx *gin.Context) {
	inp := entity.LoginForm{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest.Error())
		return
	}

	accessToken, refreshToken, err := controller.userService.SignIn(inp.Email, inp.Password)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, entity.ResponseView{
			IsSuccess: false,
			Data:      err.Error(),
		})
		return
	}

	resp := entity.ResponseView{
		IsSuccess: true,
		Data: entity.TokensResponseView{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	ctx.JSON(http.StatusOK, resp)
}

// SaveUserHandler godoc
// @Summary      Save new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 User body  entity.User true "user"
// @Success      200  {object}  entity.ResponseView{data=entity.User}
// @Router       /save-user [post]
func (controller *UserController) SaveUserHandler(ctx *gin.Context) {
	inp := entity.User{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest.Error())
		return
	}

	savedUser, err := controller.userService.SaveUser(&inp)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, entity.ResponseView{
			IsSuccess: false,
			Data:      err.Error(),
		})
		return
	}

	resp := entity.ResponseView{
		IsSuccess: true,
		Data:      savedUser,
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetUsersHandler godoc
// @Summary      Get user by params
// @Tags         User
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param 		 User body  entity.UserSearchParams true "Params for search users"
// @Success      200  {object}  entity.ResponseView{data=entity.User}
// @Router       /get-users [post]
func (controller *UserController) GetUsersHandler(ctx *gin.Context) {
	inp := entity.UserSearchParams{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest.Error())
		return
	}

	users, err := controller.userService.GetUser(ctx, &inp)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, entity.ResponseView{
			IsSuccess: false,
			Data:      err.Error(),
		})
		return
	}

	resp := entity.ResponseView{
		IsSuccess: true,
		Data:      users,
	}

	ctx.JSON(http.StatusOK, resp)
}

// RefreshTokenHandler  godoc
// @Summary      Refresh access token by refresh token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param Authorization header string true "String with the bearer started"
// @Param Refresh-Token header string true "refresh token string"
// @Success      200  {object}  entity.ResponseView{data=entity.UserSession}
// @Router       /refresh-token [get]
func (controller *UserController) RefreshTokenHandler(ctx *gin.Context) {
	authHeader := strings.Split(ctx.GetHeader("Authorization"), "Bearer")
	refreshToken := ctx.GetHeader("Refresh-Token")

	errView := entity.ResponseView{
		IsSuccess: false,
		Data:      exception.UnauthorizedError,
	}

	if len(refreshToken) == 0 || len(authHeader) < 2 {
		ctx.AbortWithStatusJSON(http.StatusOK, errView)
		return
	}

	accessToken := strings.TrimSpace(authHeader[1])

	serverConfig := config.GetServerConfig()

	tokenSubject, isValid, _ := user.ParseToken(accessToken, serverConfig.TokenKey)

	if !isValid || tokenSubject[constants.CtxUserIdKey] == nil {
		ctx.AbortWithStatusJSON(http.StatusOK, errView)
		return
	}

	accessToken, refreshToken, err := controller.userService.RefreshUserSession(ctx, refreshToken, int(tokenSubject[constants.CtxUserIdKey].(float64)))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, errView)
		return
	}

	ctx.JSON(http.StatusOK, entity.ResponseView{
		IsSuccess: true,
		Data: entity.TokensResponseView{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}
