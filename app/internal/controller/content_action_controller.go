package controller

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/service/content_action"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ContentActionController struct {
	contentActionService content_action.ServiceI
}

func NewContentActionController(contentActionService content_action.ServiceI) *ContentActionController {
	return &ContentActionController{
		contentActionService: contentActionService,
	}
}

func (controller *ContentActionController) RegisterHandlers(private, _ *gin.RouterGroup) {
	private.POST("/create_content_action", controller.CreateContentAction)
	private.DELETE("/delete_content_action/:id", controller.DeleteContentAction)
	private.GET("/get_content_actions", controller.GetContentActions)
	private.POST("/update_content_action", controller.UpdateContentAction)
}

// CreateContentAction godoc
// @Summary      Create content action
// @Security	 BearerAuth
// @Tags         ContentAction
// @Accept       json
// @Produce      json
// @Param 		 ContentAction body  entity.ContentAction true  "Only title field is accepted, the rest are ignored"
// @Success      200  {object}  entity.ResponseView{data=integer}
// @Router       /create_content_action [post]
func (controller *ContentActionController) CreateContentAction(ctx *gin.Context) {
	inp := entity.ContentAction{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest)
		return
	}

	id, err := controller.contentActionService.CreateContentAction(&inp)

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

// DeleteContentAction godoc
// @Summary      Delete content action
// @Security	 BearerAuth
// @Tags         ContentAction
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ContentAction ID"
// @Success      200  {object}  entity.ResponseView{data=object}
// @Router       /delete_content_action/{id} [delete]
func (controller *ContentActionController) DeleteContentAction(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, entity.ResponseView{
			IsSuccess: false,
			Data:      exception.BadRequest,
		})
	}

	inp := entity.ContentAction{Id: id}

	err = controller.contentActionService.DeleteContentAction(&inp)

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

// GetContentActions godoc
// @Summary      Get content actions
// @Security	 BearerAuth
// @Tags         ContentAction
// @Accept  	 json
// @Produce      json
// @Success      200  {object}  entity.ResponseView{data=[]entity.ContentAction}
// @Router       /get_content_actions [get]
func (controller *ContentActionController) GetContentActions(ctx *gin.Context) {
	list, err := controller.contentActionService.GetContentActions()

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

// UpdateContentAction godoc
// @Summary      Update content action
// @Security	 BearerAuth
// @Tags         ContentAction
// @Accept       json
// @Produce      json
// @Param 		 ContentAction body entity.ContentAction true  "Needs id and new title to change"
// @Success      200  {object}  entity.ResponseView{data=object}
// @Router       /update_content_action [post]
func (controller *ContentActionController) UpdateContentAction(ctx *gin.Context) {
	inp := entity.ContentAction{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest)
		return
	}

	err := controller.contentActionService.UpdateContentAction(&inp)

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
