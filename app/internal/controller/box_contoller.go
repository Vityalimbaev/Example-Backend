package controller

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/service/box"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type BoxController struct {
	boxService box.ServiceI
}

func NewBoxController(BoxService box.ServiceI) *BoxController {
	return &BoxController{
		boxService: BoxService,
	}
}

func (controller *BoxController) RegisterHandlers(private, _ *gin.RouterGroup) {
	private.POST("/create_box", controller.CreateBoxHandler)
	private.POST("/get_boxes", controller.GetBoxesHandler)
	private.POST("/update_box", controller.UpdateBoxHandler)
}

// CreateBoxHandler godoc
// @Summary      Save box
// @Security	 BearerAuth
// @Tags         Box
// @Accept       json
// @Produce      json
// @Param 		 Box body  entity.Box true  "id field is ignored"
// @Success      200  {object}  entity.ResponseView{data=integer}
// @Router       /create_box [post]
func (controller *BoxController) CreateBoxHandler(ctx *gin.Context) {
	inp := entity.Box{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest)
		return
	}

	id, err := controller.boxService.CreateBox(&inp)

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

// GetBoxesHandler godoc
// @Summary      Get all boxes
// @Security	 BearerAuth
// @Tags         Box
// @Accept  	 json
// @Produce      json
// @Success      200  {object}  entity.ResponseView{data=[]entity.Box}
// @Router       /get_boxes [post]
func (controller *BoxController) GetBoxesHandler(ctx *gin.Context) {
	list, err := controller.boxService.GetBoxes()

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

// UpdateBoxHandler godoc
// @Summary      Get all boxes
// @Security	 BearerAuth
// @Tags         Box
// @Accept  	 json
// @Produce      json
// @Success      200  {object}  entity.ResponseView{}
// @Router       /update_box [post]
func (controller *BoxController) UpdateBoxHandler(ctx *gin.Context) {
	inp := entity.Box{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest)
		return
	}

	err := controller.boxService.UpdateBox(&inp)

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
