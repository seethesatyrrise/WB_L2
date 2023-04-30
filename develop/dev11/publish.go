package main

import (
	"encoding/json"
	"net/http"
)

func publishError(w http.ResponseWriter, status int, err error) {
	errorJSON, err := json.Marshal(&errorModel{Err: err.Error()})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(status)
	w.Write(errorJSON)
}

func publishData(w http.ResponseWriter, status int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		publishError(w, http.StatusServiceUnavailable, err)
		return
	}
	w.WriteHeader(status)
	w.Write(response)
}
