package repository

import "github.com/madxiii/axxonsofttask/model"

type Repository interface {
	Init() map[string]model.Response
	Add(id string, data model.Response)
	Get(id string)
}
