package service

import (
	"github.com/madxiii/axxonsofttask/repository"
)

type Service struct {
	Tools *Tools
}

func New(repo repository.Repository) *Service {
	return &Service{
		Tools: NewTools(repo),
	}

}
