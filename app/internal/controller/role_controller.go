package controller

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/service/role"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type RoleController struct {
	roleService role.ServiceI
}

func NewRoleController(roleService role.ServiceI) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (controller *RoleController) RegisterHandlers(private, _ *gin.RouterGroup) {
	private.POST("/create_role", controller.CreateRole)
	private.DELETE("/delete_role/:id", controller.DeleteRole)
	private.GET("/get_roles", controller.GetRoles)
	private.POST("/update_role", controller.UpdateRole)
}

// CreateRole godoc
// @Summary      Create role
// @Security	 BearerAuth
// @Tags         Role
// @Accept       json
// @Produce      json
// @Param 		 Role body  entity.Role true  "Only title field is accepted, the rest are ignored"
// @Success      200  {object}  entity.ResponseView{data=integer}
// @Router       /create_role [post]
func (controller *RoleController) CreateRole(ctx *gin.Context) {
	inp := entity.Role{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest)
		return
	}

	id, err := controller.roleService.CreateRole(&inp)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, entity.ResponseView{
			IsSuccess: false,
			Data:      err.Error(),
		})
		return
	}

	resp := entity.ResponseView{
		IsSuccess: true,
		Data:      id,
	}

	ctx.JSON(http.StatusOK, resp)
}

// DeleteRole godoc
// @Summary      Delete role
// @Security	 BearerAuth
// @Tags         Role
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Role ID"
// @Success      200  {object}  entity.ResponseView{data=object}
// @Router       /delete_role/{id} [delete]
func (controller *RoleController) DeleteRole(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, entity.ResponseView{
			IsSuccess: false,
			Data:      exception.BadRequest,
		})
	}

	inp := entity.Role{Id: id}

	err = controller.roleService.DeleteRole(&inp)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, entity.ResponseView{
			IsSuccess: false,
			Data:      err.Error(),
		})
		return
	}

	resp := entity.ResponseView{
		IsSuccess: true,
		Data:      struct{}{},
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetRoles godoc
// @Summary      Get roles
// @Security	 BearerAuth
// @Tags         Role
// @Accept  	 json
// @Produce      json
// @Success      200  {object}  entity.ResponseView{data=[]entity.Role}
// @Router       /get_roles [get]
func (controller *RoleController) GetRoles(ctx *gin.Context) {
	list, err := controller.roleService.GetRoles()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, entity.ResponseView{
			IsSuccess: false,
			Data:      err.Error(),
		})
		return
	}

	resp := entity.ResponseView{
		IsSuccess: true,
		Data:      list,
	}

	ctx.JSON(http.StatusOK, resp)
}

// UpdateRole godoc
// @Summary      Update role
// @Security	 BearerAuth
// @Tags         Role
// @Accept       json
// @Produce      json
// @Param 		 Role body entity.Role true  "Needs id and new title to change"
// @Success      200  {object}  entity.ResponseView{data=object}
// @Router       /update_role [post]
func (controller *RoleController) UpdateRole(ctx *gin.Context) {
	inp := entity.Role{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest)
		return
	}

	err := controller.roleService.UpdateRole(&inp)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, entity.ResponseView{
			IsSuccess: false,
			Data:      err.Error(),
		})
		return
	}

	resp := entity.ResponseView{
		IsSuccess: true,
		Data:      struct{}{},
	}

	ctx.JSON(http.StatusOK, resp)
}
