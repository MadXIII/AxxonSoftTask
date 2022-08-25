package http

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) Request(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		param := r.FormValue("id")
		res, err := h.service.Tools.SendRequestID(param)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		slice, err := json.Marshal(res)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		okResponse(w, http.StatusOK, slice)
	default:
		wrongMethodResponse(w, http.StatusMethodNotAllowed, wrongMehtod)
	}
}
