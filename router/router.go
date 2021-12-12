package router

import (
	"github.com/ruansheng/fly"
	"github.com/ruansheng/fly_demo/controller"
	"github.com/ruansheng/fly_demo/middleware"
)

func AddRouter(fly *fly.Fly) {
	fly.GET("/abc", controller.Abc)

	group := fly.Group("/group")
	group.Use(middleware.IpLimit())
	group.GET("/def", controller.Def, middleware.MustLogin)
}
