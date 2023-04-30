package main

import (
	"errors"
	"net/http"
)

func parseRequest(req *http.Request) (*event, error) {
	keys := req.URL.Query()
	e := &event{}

	if keys["date"] != nil {
		e.date = keys["date"][0]
		if !validateDay(e.date) {
			return nil, errors.New("wrong date string")
		}
	}
	if keys["month"] != nil {
		e.date = keys["month"][0]
		if !validateMonth(e.date) {
			return nil, errors.New("wrong month string")
		}
	}
	if keys["textOld"] != nil {
		e.textOld = keys["textOld"][0]
	}
	if keys["text"] != nil {
		e.text = keys["text"][0]
	}
	return e, nil
}
