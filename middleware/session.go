package middleware

import (
	"fmt"
	"github.com/ruansheng/fly"
)

func MustLogin(next fly.HandlerFunc) fly.HandlerFunc {
	return func(ctx fly.Context) error {
		fmt.Println("middleware->MustLogin", ctx.Request().UserAgent())
		return next(ctx)
	}
}
