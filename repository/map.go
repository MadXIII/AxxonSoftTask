package repository

import (
	"errors"
	"fmt"

	"github.com/madxiii/axxonsofttask/model"
)

var requestNotFound error = errors.New("request is not found")
var responseNotFound error = errors.New("response is not found")

type HashMap struct {
	reqMap map[string]model.Request
	resMap map[string]model.Response
}

func NewHashMap() *HashMap {
	return &HashMap{
		reqMap: make(map[string]model.Request),
		resMap: make(map[string]model.Response),
	}
}
func (h *HashMap) AddRequest(id string, data model.Request) {
	h.reqMap[id] = data
	fmt.Printf("Request is stored in map by id: %s\n", id)
}
func (h *HashMap) AddResponse(data model.Response) {
	h.resMap[data.ID] = data
	fmt.Printf("Response is stored in map by id: %s\n", data.ID)
}

func (h *HashMap) GetRequest(id string) (model.Request, error) {
	var res model.Request
	for key, val := range h.reqMap {
		if id == key {
			return val, nil
		}
	}
	return res, requestNotFound
}

func (h *HashMap) GetResponse(id string) (model.Response, error) {
	var res model.Response
	for key, val := range h.resMap {
		if id == key {
			return val, nil
		}
	}
	return res, responseNotFound
}
