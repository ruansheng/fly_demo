package controller

import (
	"fmt"
	"github.com/ruansheng/fly"
	"net/http"
)

func Abc(ctx fly.Context) error {
	fmt.Println("controller->Abc")
	hello := fmt.Sprintf("hello abc %s fly!", ctx.Path())
	return ctx.JSON(http.StatusOK, hello)
}
