package middleware

import (
	"fmt"
	"github.com/ruansheng/fly"
)

func CheckToken() fly.MiddlewareFunc {
	return func(next fly.HandlerFunc) fly.HandlerFunc {
		return func(ctx fly.Context) error {
			fmt.Println("middleware->CheckToken", ctx.RealIp())
			return next(ctx)
		}
	}
}
