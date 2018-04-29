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
	tagController := &controller.TagController{Base: baseController}
	articleController := &controller.ArticleController{Base: baseController}
	evaluatoinController := &controller.EvaluationController{Base: baseController}

	e.GET("/", indexController.Default)

	apiGroup := e.Group("/api")

	/* 验证接口 */
	apiGroup.POST("/auth/login", authController.Login)

	/* 附件接口 */
	apiGroup.POST("/attachment/upload", attachmentController.Upload, authMiddleware.AuthToken)

	/* 分类接口 */
	apiGroup.GET("/categories", categoryController.GetCategories)
	apiGroup.GET("/category/:id", categoryController.GetCategory)
	apiGroup.POST("/category", categoryController.AddCategory, authMiddleware.AuthToken)
	apiGroup.POST("/category/:id", categoryController.UpdateCategory, authMiddleware.AuthToken)
	apiGroup.POST("/category/:id/delete", categoryController.DelCategory, authMiddleware.AuthToken)

	/* 标签接口 */
	apiGroup.GET("/tags", tagController.GetTags)
	apiGroup.GET("/tag/:id", tagController.GetTag)
	apiGroup.POST("/tag", tagController.AddTag, authMiddleware.AuthToken)
	apiGroup.POST("/tag/:id", tagController.UpdateTag, authMiddleware.AuthToken)
	apiGroup.POST("/tag/:id/delete", tagController.DelTag, authMiddleware.AuthToken)

	/* 文章接口 */
	apiGroup.POST("/article", articleController.AddArticle, authMiddleware.AuthToken)
	apiGroup.GET("/articles", articleController.GetArticles)
	apiGroup.GET("/article/:id", articleController.GetArticle)
	apiGroup.GET("/article/:id/:flag", articleController.GetArticlesByFlagWithId)
	apiGroup.POST("/article/:id/evaluate", evaluatoinController.AddEvaluation, authMiddleware.AuthToken)
	apiGroup.POST("/article/:id/delete", articleController.DelArticle, authMiddleware.AuthToken)

	/* 评论接口 */
	apiGroup.GET("/evaluations/:id", evaluatoinController.GetEvaluations)
	apiGroup.GET("/evaluation/:id", evaluatoinController.GetEvaluation)
	apiGroup.POST("/evaluation/:id/delete", evaluatoinController.DelEvaluation, authMiddleware.AuthToken)

}
