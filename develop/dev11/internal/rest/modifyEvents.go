package rest

import (
	"WB-L2/develop/dev11/internal/models"
	"WB-L2/develop/dev11/internal/utils"
	"errors"
	"net/http"
)

func (c *Calendar) CreateEvent(w http.ResponseWriter, req *http.Request) {
	c.modifyEvent(w, req, "create", &eventNeeds{date: true, text: true})
}

func (c *Calendar) UpdateEvent(w http.ResponseWriter, req *http.Request) {
	c.modifyEvent(w, req, "update", &eventNeeds{date: true, text: true, textOld: true})
}

func (c *Calendar) DeleteEvent(w http.ResponseWriter, req *http.Request) {
	c.modifyEvent(w, req, "delete", &eventNeeds{date: true, text: true})
}

func (c *Calendar) modifyEvent(w http.ResponseWriter, req *http.Request, modify string, needs *eventNeeds) {
	keys := req.URL.Query()
	e, err := ParseRequest(keys, needs)
	if err != nil {
		utils.PublishError(w, http.StatusBadRequest, err)
		return
	}

	var response string
	var status int

	switch modify {
	case "create":
		response = c.create(e)
		status = http.StatusOK
	case "update":
		response, status, err = c.update(e)
	case "delete":
		response, status, err = c.delete(e)
	default:
		status = http.StatusServiceUnavailable
		err = errors.New("something weird")
	}
	if err != nil {
		utils.PublishError(w, status, err)
		return
	}

	utils.PublishData(w, status, &models.ResponseModel{Resp: response})
}

func (c *Calendar) create(e *event) string {
	if c.days[e.date] == nil {
		c.days[e.date] = make(map[string]struct{})
	}
	c.days[e.date][e.text] = struct{}{}
	response := "event " + e.text + " was created"

	return response
}

func (c *Calendar) update(e *event) (string, int, error) {
	_, ok := c.days[e.date][e.textOld]
	if !ok {
		return "", http.StatusBadRequest, errors.New("can't find event")
	}

	delete(c.days[e.date], e.textOld)
	c.days[e.date][e.text] = struct{}{}
	response := "event " + e.textOld + " was updated to " + e.text

	return response, http.StatusOK, nil
}

func (c *Calendar) delete(e *event) (string, int, error) {
	_, ok := c.days[e.date][e.text]
	if !ok {
		return "", http.StatusBadRequest, errors.New("can't find event")
	}

	delete(c.days[e.date], e.text)

	response := "event " + e.text + " was deleted"

	return response, http.StatusOK, nil
}
