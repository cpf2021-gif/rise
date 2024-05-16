package server

import (
	"github.com/cpf2021-gif/rise/common/net/chttp"
	"github.com/cpf2021-gif/rise/internal/conf"
	"github.com/cpf2021-gif/rise/internal/handler"
	"github.com/cpf2021-gif/rise/internal/service"
)

func NewHTTP(conf *conf.Conf) *chttp.Server {
	s := chttp.NewServer(conf.Server)
	svc := service.NewSerive(conf)

	handler.InitRouter(s, svc)

	err := s.Start()
	if err != nil {
		panic(err)
	}

	return s
}
