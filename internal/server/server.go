package server

import (
	"context"

	"github.com/cpf2021-gif/rise/common/net/chttp"
	"github.com/cpf2021-gif/rise/internal/handler"
	"github.com/cpf2021-gif/rise/internal/service"
)

type App struct {
	server *chttp.Server
}

func NewApp(s *chttp.Server, svc *service.Service) *App {
	handler.InitRouter(s, svc)

	err := s.Start()
	if err != nil {
		panic(err)
	}
	return &App{
		server: s,
	}
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
