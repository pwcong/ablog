package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pwcong/ablog/config"
	"github.com/pwcong/ablog/controller"
	"github.com/pwcong/ablog/middleware"
	"github.com/pwcong/ablog/service"
)

func Init(e *echo.Echo, conf *config.Config, db *gorm.DB) {

	e.Static("/", "view/dist")
	e.Static("/public", "public")
	e.Static("/doc", "doc")

	authMiddleware := middleware.AuthMiddleware{Conf: conf}

	baseService := &service.BaseService{Conf: conf, DB: db}
	baseController := &controller.BaseController{Conf: conf, Service: baseService}

	indexController := &controller.IndexController{Base: baseController}
	authController := &controller.AuthController{Base: baseController}
	attachmentController := &controller.AttachmentController{Base: baseController}
	categoryController := &controller.CategoryController{Base: baseController}

	e.GET("/", indexController.Default)

	apiGroup := e.Group("/api")
	apiGroup.POST("/auth/login", authController.Login)
	apiGroup.POST("/attachment/upload", attachmentController.Upload, authMiddleware.AuthToken)
	apiGroup.GET("/categories", categoryController.GetCategories)

}
