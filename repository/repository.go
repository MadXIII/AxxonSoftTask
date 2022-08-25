package repository

import "github.com/madxiii/axxonsofttask/model"

type Store interface {
	AddRequest(id string, data model.Request)
	AddResponse(data model.Response)
	GetRequest(id string) (model.Request, error)
	GetResponse(id string) (model.Response, error)
}

type Repository struct {
	Store
}

func New(store Store) *Repository {
	return &Repository{
		Store: store,
	}
}
