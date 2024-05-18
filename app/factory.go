package app

import (
	"github.com/cpf2021-gif/rise/common/cconf"
	"github.com/cpf2021-gif/rise/common/database"
	"github.com/cpf2021-gif/rise/common/net/chttp"
	"github.com/cpf2021-gif/rise/internal/repo"
	"github.com/cpf2021-gif/rise/internal/repo/db"
	"github.com/cpf2021-gif/rise/internal/server"
	"github.com/cpf2021-gif/rise/internal/service"

	"go.uber.org/dig"
)

var container = dig.New()

func init() {
	// config
	_ = container.Provide(cconf.NewConfig)

	// database.DB
	_ = container.Provide(database.NewDB, dig.As(new(db.DBTX)))

	// db.Queries
	_ = container.Provide(db.New)

	// repo
	_ = container.Provide(repo.NewRepo)

	// service
	_ = container.Provide(service.NewSerive)

	// server
	_ = container.Provide(chttp.NewServer)

	// app
	_ = container.Provide(server.NewApp)
}

func ConstructServer() (*server.App, error) {
	var app *server.App
	err := container.Invoke(func(a *server.App) {
		app = a
	})

	return app, err
}
