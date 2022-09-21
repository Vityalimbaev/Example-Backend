package controller

import (
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/service/record"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type RecordController struct {
	recordService record.ServiceI
}

func NewRecordController(recordService record.ServiceI) *RecordController {
	return &RecordController{
		recordService: recordService,
	}
}

func (controller *RecordController) RegisterHandlers(private, _ *gin.RouterGroup) {
	private.POST("/create_record", controller.CreateRecord)
	private.POST("/get_records", controller.GetRecords)
	private.POST("/update_record", controller.UpdateRecord)
}

// CreateRecord godoc
// @Summary      Create record
// @Security	 BearerAuth
// @Tags         Record
// @Accept       json
// @Produce      json
// @Param 		 Record body  entity.Record true  "pcode, content_state_id are required; archived_date, last_treat are optional"
// @Success      200  {object}  entity.Record{data=integer}
// @Router       /create_record [post]
func (controller *RecordController) CreateRecord(ctx *gin.Context) {
	inp := entity.Record{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest)
		return
	}

	id, err := controller.recordService.CreateRecord(&inp)

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

// GetRecords godoc
// @Summary      Get records
// @Security	 BearerAuth
// @Tags         Record
// @Accept  	 json
// @Produce      json
// @Param        RecordSearchParams body entity.RecordSearchParams true  "The fields EndCreationDate, EndArchivedDate, EndLastTreat default to the current date "
// @Success      200  {object}  entity.ResponseView{data=[]entity.Record}
// @Router       /get_records [post]
func (controller *RecordController) GetRecords(ctx *gin.Context) {
	inp := entity.RecordSearchParams{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(http.StatusOK, exception.BadRequest)
		return
	}

	list, err := controller.recordService.GetRecords(&inp)

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

// UpdateRecord godoc
// @Summary      Update record
// @Security	 BearerAuth
// @Tags         Record
// @Accept       json
// @Produce      json
// @Param 		 Record body entity.Record true  "id is required"
// @Success      200  {object}  entity.ResponseView{data=object}
// @Router       /update_record [post]
func (controller *RecordController) UpdateRecord(ctx *gin.Context) {
	inp := entity.Record{}

	if err := ctx.BindJSON(&inp); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(exception.GetHttpStatusCode(exception.BadRequest), exception.BadRequest)
		return
	}

	err := controller.recordService.UpdateRecord(&inp)

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
