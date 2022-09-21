package controller

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/service/content_state"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ContentStateController struct {
	contentStateService content_state.ServiceI
}

func NewContentStateController(contentStateService content_state.ServiceI) *ContentStateController {
	return &ContentStateController{
		contentStateService: contentStateService,
	}
}

func (controller *ContentStateController) RegisterHandlers(private, _ *gin.RouterGroup) {
	private.POST("/create_content_state", controller.CreateContentState)
	private.DELETE("/delete_content_state/:id", controller.DeleteContentState)
	private.GET("/get_content_states", controller.GetContentStates)
	private.POST("/update_content_state", controller.UpdateContentState)
}

// CreateContentState godoc
// @Summary      Create content state
// @Security	 BearerAuth
// @Tags         ContentState
// @Accept       json
// @Produce      json
// @Param 		 ContentState body  entity.ContentState true  "Only title field is accepted, the rest are ignored"
// @Success      200  {object}  entity.ResponseView{data=integer}
// @Router       /create_content_state [post]
func (controller *ContentStateController) CreateContentState(ctx *gin.Context) {
	inp := entity.ContentState{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest)
		return
	}

	id, err := controller.contentStateService.CreateContentState(&inp)

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

// DeleteContentState godoc
// @Summary      Delete content state
// @Security	 BearerAuth
// @Tags         ContentState
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ContentState ID"
// @Success      200  {object}  entity.ResponseView{data=object}
// @Router       /delete_content_state/{id} [delete]
func (controller *ContentStateController) DeleteContentState(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, entity.ResponseView{
			IsSuccess: false,
			Data:      exception.BadRequest,
		})
	}

	inp := entity.ContentState{Id: id}

	err = controller.contentStateService.DeleteContentState(&inp)

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

// GetContentStates godoc
// @Summary      Get content states
// @Security	 BearerAuth
// @Tags         ContentState
// @Accept  	 json
// @Produce      json
// @Success      200  {object}  entity.ResponseView{data=[]entity.ContentState}
// @Router       /get_content_states [get]
func (controller *ContentStateController) GetContentStates(ctx *gin.Context) {
	list, err := controller.contentStateService.GetContentStates()

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

// UpdateContentState godoc
// @Summary      Update content state
// @Security	 BearerAuth
// @Tags         ContentState
// @Accept       json
// @Produce      json
// @Param 		 ContentState body entity.ContentState true  "Needs id and new title to change"
// @Success      200  {object}  entity.ResponseView{data=object}
// @Router       /update_content_state [post]
func (controller *ContentStateController) UpdateContentState(ctx *gin.Context) {
	inp := entity.ContentState{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest)
		return
	}

	err := controller.contentStateService.UpdateContentState(&inp)

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
