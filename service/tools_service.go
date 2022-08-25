package service

import (
	"errors"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/madxiii/axxonsofttask/model"
	"github.com/madxiii/axxonsofttask/repository"
)

var errID error = errors.New("wrong ID")

type Tools struct {
	repo repository.Store
}

func NewTools(repo repository.Repository) *Tools {
	return &Tools{
		repo: repo,
	}
}

func (t *Tools) PrepareRequestToStore(inp model.Request) error {
	uinqID, err := generateID()
	if err != nil {
		return err
	}

	t.repo.AddRequest(uinqID, inp)

	return nil
}

func (t *Tools) PrepareResponseToStore(resp *http.Response) (model.Response, error) {
	uinqID, err := generateID()
	if err != nil {
		return model.Response{}, err
	}

	filled := fillResponse(uinqID, resp)

	t.repo.AddResponse(filled)

	return filled, nil
}

func (t *Tools) SendRequestID(id string) (model.Request, error) {
	if len(id) != 36 {
		return model.Request{}, errID
	}
	return t.repo.GetRequest(id)
}

func (t *Tools) SendResponseID(id string) (model.Response, error) {
	if len(id) != 36 {
		return model.Response{}, errID
	}
	return t.repo.GetResponse(id)
}

func generateID() (string, error) {
	hash, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return hash.String(), nil
}

func fillResponse(id string, resp *http.Response) model.Response {
	var res model.Response
	res.ID = id
	res.Status = resp.Status
	res.Length = resp.ContentLength
	res.Headers = make(map[string][]string)

	for key, val := range resp.Header {
		res.Headers[key] = val
	}

	return res
}
