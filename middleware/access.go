package middleware

import (
	"fmt"
	"github.com/ruansheng/fly"
)

func Access() fly.MiddlewareFunc {
	return func(next fly.HandlerFunc) fly.HandlerFunc {
		return func(ctx fly.Context) error {
			fmt.Println("middleware->Access:", ctx.Path())
			return next(ctx)
		}
	}
}
