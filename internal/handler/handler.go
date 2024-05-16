package handler

import (
	"github.com/cpf2021-gif/rise/common/net/chttp"
	"github.com/cpf2021-gif/rise/internal/service"
)

var svc *service.Service

func InitRouter(s *chttp.Server, service *service.Service) {
	svc = service

	g := s.Group("/v1")
	{
		g.GET("/ping", echo)
	}

	ug := g.Group("/user")
	{
		ug.POST("/create", userCreate)
		ug.GET("/datail", userDetail)
		ug.GET("/list", userList)
	}
}
