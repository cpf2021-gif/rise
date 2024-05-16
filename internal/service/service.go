package service

import (
	"github.com/cpf2021-gif/rise/internal/conf"
	"github.com/cpf2021-gif/rise/internal/repo"
)

type Service struct {
	conf *conf.Conf
	repo *repo.Repo
}

func NewSerive(conf *conf.Conf) *Service {
	return &Service{
		conf: conf,
		repo: repo.NewRepo(conf),
	}
}
