package main

import (
	"errors"
	"net/http"
)

func (c *calendar) modifyEvent(w http.ResponseWriter, req *http.Request) {
	e, err := parseRequest(req)
	if err != nil {
		publishError(w, http.StatusBadRequest, err)
		return
	}
	var response string
	var status int
	switch req.URL.Path[1:] {
	case "create_event":
		status = c.create(e)
		response = "event " + e.text + " was created"
	case "delete_event":
		status = c.delete(e)
		response = "event " + e.text + " was deleted"
	case "update_event":
		status = c.update(e)
		response = "event " + e.textOld + " was updated to " + e.text
	default:
		publishError(w, http.StatusServiceUnavailable, errors.New("???"))
		return
	}
	if err != nil {
		publishError(w, status, err)
		return
	}

	publishData(w, status, &responseModel{Resp: response})
}

func (c *calendar) create(e *event) int {
	if e.date == "" || e.text == "" {
		return http.StatusBadRequest
	}
	if c.days[e.date] == nil {
		c.days[e.date] = make(map[string]struct{})
	}
	c.days[e.date][e.text] = struct{}{}
	return http.StatusOK
}

func (c *calendar) update(e *event) int {
	if e.date == "" || e.text == "" || e.textOld == "" {
		return http.StatusBadRequest
	}
	_, ok := c.days[e.date][e.textOld]
	if !ok {
		return http.StatusBadRequest
	}
	delete(c.days[e.date], e.textOld)
	c.days[e.date][e.text] = struct{}{}
	return http.StatusOK
}

func (c *calendar) delete(e *event) int {
	if e.date == "" || e.text == "" {
		return http.StatusBadRequest
	}
	_, ok := c.days[e.date][e.text]
	if !ok {
		return http.StatusBadRequest
	}
	delete(c.days[e.date], e.text)
	return http.StatusOK
}
