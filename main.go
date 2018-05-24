package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pwcong/ablog/config"
	"github.com/pwcong/ablog/controller"
	"github.com/pwcong/ablog/db"
	"github.com/pwcong/ablog/middleware"
	"github.com/pwcong/ablog/model"
	"github.com/pwcong/ablog/router"
	"github.com/pwcong/ablog/utils"
)

func initMiddlewares(e *echo.Echo, conf *config.Config) {
	middleware.Init(e, conf)
}

func initRoutes(e *echo.Echo, conf *config.Config, db *gorm.DB) {
	router.Init(e, conf, db)
}

func initDB(db *gorm.DB) {
	db.AutoMigrate(
		&model.Attachment{},
		&model.Log{},
		&model.Article{},
		&model.Tag{},
		&model.Category{},
		&model.Evaluation{},
	)
}

func main() {

	// 初始化配置
	conf, err := config.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	mySQLConfig, ok := conf.Databases["mysql"]
	if !ok {
		log.Fatal("Can not load configuration of MySQL")
	}

	orm := db.ORM{DB: nil, Name: "mysql"}

	orm.Open(
		mySQLConfig.Username,
		mySQLConfig.Password,
		mySQLConfig.Host+":"+strconv.Itoa(mySQLConfig.Port),
		mySQLConfig.DBName)

	defer orm.Close()

	initDB(orm.DB)

	e := echo.New()

	// 全局错误处理
	e.HTTPErrorHandler = func(err error, c echo.Context) {

		e.Logger.Error(err)

		c.JSON(http.StatusOK, controller.BaseResponseJSON{
			Success: false,
			Code:    controller.STATUS_ERROR,
			Message: err.Error() + ", url=" + c.Request().RequestURI,
		})
	}

	// 初始化中间件
	initMiddlewares(e, &conf)
	// 初始化路由
	initRoutes(e, &conf, orm.DB)

	// 运行服务
	if conf.Server.Port == 80 {
		e.Logger.Fatal(e.Start(conf.Server.Host))
	} else {
		e.Logger.Fatal(e.Start(conf.Server.Host + ":" + strconv.Itoa(conf.Server.Port)))
	}

}

func init() {
	// 初始化目录
	err := utils.MkdirIFNotExist("public")
	if err != nil {
		log.Fatal(err)
	}

	err = utils.MkdirIFNotExist("view")
	if err != nil {
		log.Fatal(err)
	}

}
