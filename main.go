package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/cpf2021-gif/rise/app"
)

func main() {
	app, err := app.ConstructServer()
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for {
		sig := <-ch
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			_ = app.Shutdown(context.Background())
			return
		default:
			return
		}
	}
}
