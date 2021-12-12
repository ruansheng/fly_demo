package middleware

import (
	"fmt"
	"github.com/ruansheng/fly"
)

func IpLimit() fly.MiddlewareFunc {
	return func(next fly.HandlerFunc) fly.HandlerFunc {
		return func(ctx fly.Context) error {
			fmt.Println("middleware->IpLimit", ctx.RealIp())
			return next(ctx)
		}
	}
}
