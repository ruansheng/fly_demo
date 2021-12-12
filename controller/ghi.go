package controller

import (
	"fmt"
	"github.com/ruansheng/fly"
	"net/http"
)

func Ghi(ctx fly.Context) error {
	fmt.Println("controller->Ghi")
	hello := fmt.Sprintf("hello ghi %s fly!", ctx.Path())
	return ctx.JSON(http.StatusOK, hello)
}
