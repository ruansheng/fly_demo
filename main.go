package main

import (
	"github.com/ruansheng/fly"
	"github.com/ruansheng/fly/middleware"
	selfMiddleware "github.com/ruansheng/fly_demo/middleware"
	"github.com/ruansheng/fly_demo/router"
	"net/http"
)

func main() {
	f := fly.NewFly(true)

	f.Pre(selfMiddleware.Access())
	f.Use(middleware.Recover())

	router.AddRouter(f)

	server := &http.Server{
		Addr:    ":8080",
		Handler: f,
	}

	f.StartServer(server)
}
