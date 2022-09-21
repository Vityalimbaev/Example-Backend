package adapter

import (
	"github.com/Vityalimbaev/Example-Backend/config"
	"github.com/Vityalimbaev/Example-Backend/docs"
	"github.com/Vityalimbaev/Example-Backend/internal/adapter/database"
	"github.com/Vityalimbaev/Example-Backend/internal/adapter/server"
	"github.com/Vityalimbaev/Example-Backend/internal/controller"
	boxRepo "github.com/Vityalimbaev/Example-Backend/internal/repository/box"
	contentActionRepo "github.com/Vityalimbaev/Example-Backend/internal/repository/content_action"
	contentStateRepo "github.com/Vityalimbaev/Example-Backend/internal/repository/content_state"
	historyRepo "github.com/Vityalimbaev/Example-Backend/internal/repository/history"
	recordRepo "github.com/Vityalimbaev/Example-Backend/internal/repository/record"
	roleRepo "github.com/Vityalimbaev/Example-Backend/internal/repository/role"
	userRepo "github.com/Vityalimbaev/Example-Backend/internal/repository/user"
	"github.com/Vityalimbaev/Example-Backend/internal/service/history"

	"github.com/Vityalimbaev/Example-Backend/internal/service/box"
	"github.com/Vityalimbaev/Example-Backend/internal/service/content_action"
	"github.com/Vityalimbaev/Example-Backend/internal/service/content_state"
	"github.com/Vityalimbaev/Example-Backend/internal/service/record"
	"github.com/Vityalimbaev/Example-Backend/internal/service/role"
	"github.com/Vityalimbaev/Example-Backend/internal/service/user"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type ServicePool struct {
	UserService          user.ServiceI
	ContentActionService content_action.ServiceI
	ContentStateService  content_state.ServiceI
	RoleService          role.ServiceI
	RecordService        record.ServiceI
	BoxService           box.ServiceI
	HistoryService       history.ServiceI
}

func InitApp() {
	db := database.GetDbConnection(config.GetDbConfig())
	database.UpDBMigrations(db)

	router := server.GetRouter()
	servicePool := InitServices(db)

	InitHandlers(router, servicePool)

	SetUpSwagger(router)

	_ = server.Run(router)
}

func InitServices(db *sqlx.DB) *ServicePool {
	return &ServicePool{
		UserService:          user.NewService(userRepo.NewRepository(db)),
		RoleService:          role.NewService(roleRepo.NewRepository(db)),
		ContentActionService: content_action.NewService(contentActionRepo.NewRepository(db)),
		ContentStateService:  content_state.NewService(contentStateRepo.NewRepository(db)),
		RecordService:        record.NewService(recordRepo.NewRepository(db)),
		BoxService:           box.NewService(boxRepo.NewRepository(db)),
		HistoryService:       history.NewService(historyRepo.NewRepository(db)),
	}
}

func InitHandlers(router *gin.Engine, pool *ServicePool) {

	publicGroup := router.Group("/api")
	privateGroup := router.Group("/api")
	privateGroup.Use(controller.AuthMiddleware)

	userController := controller.NewUserController(pool.UserService)
	userController.RegisterHandlers(privateGroup, publicGroup)

	roleController := controller.NewRoleController(pool.RoleService)
	roleController.RegisterHandlers(privateGroup, publicGroup)

	contentActionController := controller.NewContentActionController(pool.ContentActionService)
	contentActionController.RegisterHandlers(privateGroup, publicGroup)

	contentStateController := controller.NewContentStateController(pool.ContentStateService)
	contentStateController.RegisterHandlers(privateGroup, publicGroup)

	recordController := controller.NewRecordController(pool.RecordService)
	recordController.RegisterHandlers(privateGroup, publicGroup)

	boxController := controller.NewBoxController(pool.BoxService)
	boxController.RegisterHandlers(privateGroup, publicGroup)

	historyController := controller.NewHistoryController(pool.HistoryService)
	historyController.RegisterHandlers(privateGroup, publicGroup)
}

// @title Sadko Archive API
// @version 1.0
// @description Sadko Archive REST API.

// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func SetUpSwagger(router *gin.Engine) {
	serverConf := config.GetServerConfig()

	docs.SwaggerInfo.Host = serverConf.ExternalDomainName + ":" + serverConf.ExternalPort
	docs.SwaggerInfo.Schemes = []string{"http"}

	url := ginSwagger.URL(serverConf.Schema + "://" + serverConf.ExternalDomainName + ":" + serverConf.ExternalPort + "/swagger/doc.json")

	logrus.Debug(serverConf.Schema, serverConf.ExternalDomainName, serverConf.ExternalPort)
	logrus.Debug(url)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
