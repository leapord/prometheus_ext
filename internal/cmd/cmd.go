package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/leapord/prometheusx/internal/controller"
	"github.com/leapord/prometheusx/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetIndexFolder(true)
			s.SetServerRoot("resource/public/html")
			// 无需权限
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					middleware.GlobalExceptionMiddleware,
				)
				group.Bind(
					controller.Hello,
					controller.Authentication,
					controller.Target,
				)
			})
			// 权限验证
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					middleware.TokenMiddleware,
					middleware.GlobalExceptionMiddleware,
				)
				group.Bind(
					controller.Group,
					controller.Node,
					controller.User,
				)
			})
			s.Run()
			return nil
		},
	}
)
