package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/suzujun/yatteiki-cloud/goapp/config"
	"github.com/suzujun/yatteiki-cloud/goapp/controller"
	"github.com/suzujun/yatteiki-cloud/goapp/dao"
)

func main() {

	setupDb()
	todoDao := dao.NewTodoDao()

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		ctrl := controller.Ping{}
		v1.GET("/ping", ctrl.Get)
		v1.GET("/pingdb", ctrl.GetDb)

		todos := v1.Group("/todos")
		{
			ctrl := controller.NewTodo(todoDao)
			todos.GET("", ctrl.GetList)
			todos.POST("", ctrl.Post)
			todos.GET("/:id", ctrl.Get)
			todos.PUT("/:id", ctrl.Put)
			todos.DELETE("/:id", ctrl.Delete)
		}
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}

func setupDb() {

	if err := config.Init(); err != nil {
		panic(errors.Wrap(err, "config.Init failed"))
	}

	conf := config.Get()

	dao.Initialize(conf.DbMasterConfig(), conf.DbSlaveConfig())
}
