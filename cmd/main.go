package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/cpf2021-gif/rise/common/cconf"
	"github.com/cpf2021-gif/rise/internal/conf"
	"github.com/cpf2021-gif/rise/internal/server"
)

var filePath = flag.String("conf", "etc/config.yaml", "config file path")

func main() {
	flag.Parse()

	c := new(conf.Conf)
	if err := cconf.Unmarshal(*filePath, c); err != nil {
		panic(err)
	}
	srv := server.NewHTTP(c)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for {
		sig := <-ch
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			_ = srv.Shutdown(context.Background())
			return
		default:
			return
		}
	}
}
