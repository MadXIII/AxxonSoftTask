package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/madxiii/axxonsofttask/model"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var body model.Request

		slice, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if err = json.Unmarshal(slice, &body); err != nil {
			errorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if err = h.service.Tools.PrepareRequestToStore(body); err != nil {
			errorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		req, err := http.NewRequest(body.Method, body.URL, nil)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, err := h.service.Tools.PrepareResponseToStore(resp)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		slice, err = json.Marshal(res)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		okResponse(w, http.StatusOK, slice)
	default:
		wrongMethodResponse(w, http.StatusMethodNotAllowed, wrongMehtod)
	}
}
