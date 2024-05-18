package service

import (
	"github.com/cpf2021-gif/rise/internal/repo"
)

type Service struct {
	repo *repo.Repo
}

func NewSerive(repo *repo.Repo) *Service {
	return &Service{
		repo: repo,
	}
}
