package controller

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/service/history"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type HistoryController struct {
	HistoryService history.ServiceI
}

func NewHistoryController(HistoryService history.ServiceI) *HistoryController {
	return &HistoryController{
		HistoryService: HistoryService,
	}
}

func (controller *HistoryController) RegisterHandlers(private, _ *gin.RouterGroup) {
	private.POST("/create_History", controller.CreateActionHistory)
	private.POST("/get_action_history", controller.GetActionHistory)
}

func (controller *HistoryController) CreateActionHistory(ctx *gin.Context) {
	inp := entity.ActionHistory{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest)
		return
	}

	id, err := controller.HistoryService.CreateActionHistory(&inp)

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

func (controller *HistoryController) GetActionHistory(ctx *gin.Context) {
	list, err := controller.HistoryService.GetActionHistory()

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
