package app

import (
	"time"

	"github.com/BeanWei/freq/internal/app/hello"
	"github.com/BeanWei/freq/internal/g"
	"github.com/BeanWei/freq/internal/http/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func NewHTTPServer() {
	svr := server.Default(
		server.WithHostPorts(g.Cfg().Server.Address),
		server.WithExitWaitTime(time.Second*0), // Only for debug
	)

	// 注册全局中间件
	svr.Use(
		middleware.ErrorHandler(),
	)

	hello.Rgister(svr)

	svr.Spin()
}
