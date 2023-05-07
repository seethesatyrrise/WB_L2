package utils

import (
	"WB-L2/develop/dev11/internal/models"
	"encoding/json"
	"net/http"
)

func PublishError(w http.ResponseWriter, status int, err error) {
	errorJSON, err := json.Marshal(&models.ErrorModel{Err: err.Error()})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(status)
	w.Write(errorJSON)
}

func PublishData(w http.ResponseWriter, status int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		PublishError(w, http.StatusServiceUnavailable, err)
		return
	}
	w.WriteHeader(status)
	w.Write(response)
}
