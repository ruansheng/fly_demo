package controller

import (
	"fmt"
	"github.com/ruansheng/fly"
	"net/http"
)

func Def(ctx fly.Context) error {
	fmt.Println("controller->Def")
	hello := fmt.Sprintf("hello def %s fly!", ctx.Path())
	return ctx.JSON(http.StatusOK, hello)
}
