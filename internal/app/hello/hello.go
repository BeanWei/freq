package hello

import (
	"github.com/BeanWei/freq/internal/app/hello/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Rgister(s *server.Hertz) {
	s.GET("/say", api.Say)
}
