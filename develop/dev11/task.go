package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type calendar struct {
	days map[string]map[string]struct{}
}

type event struct {
	date    string
	text    string
	textOld string
}

type getEventsModel struct {
	Result []DayEventsModel `json:"result"`
}

type DayEventsModel struct {
	Date   string `json:"date"`
	Events string `json:"events"`
}

func main() {
	c := &calendar{days: make(map[string]map[string]struct{})}

	http.HandleFunc("/create_event", c.modifyEvent)
	http.HandleFunc("/update_event", c.modifyEvent)
	http.HandleFunc("/delete_event", c.modifyEvent)
	http.HandleFunc("/events_for_day", c.getEvents)
	http.HandleFunc("/events_for_week", c.getEvents)
	http.HandleFunc("/events_for_month", c.getEvents)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseRequest(req *http.Request) *event {
	keys := req.URL.Query()
	e := &event{}

	if keys["date"] != nil {
		e.date = keys["date"][0]
	}
	if keys["month"] != nil {
		e.date = keys["month"][0]
	}
	if keys["textOld"] != nil {
		e.textOld = keys["textOld"][0]
	}
	if keys["text"] != nil {
		e.text = keys["text"][0]
	}
	return e
}

func (c *calendar) modifyEvent(w http.ResponseWriter, req *http.Request) {
	e := parseRequest(req)
	var err int
	var response string
	switch req.URL.Path[1:] {
	case "create_event":
		err = c.create(e)
		response = "event " + e.text + " was created\n"
	case "delete_event":
		err = c.delete(e)
		response = "event " + e.text + " was deleted\n"
	case "update_event":
		err = c.update(e)
		response = "event " + e.textOld + " was updated to " + e.text + "\n"
	default:
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	if err != 0 {
		w.WriteHeader(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, response)
}

func (c *calendar) create(e *event) int {
	if e.date == "" || e.text == "" {
		return http.StatusBadRequest
	}
	if c.days[e.date] == nil {
		c.days[e.date] = make(map[string]struct{})
	}
	c.days[e.date][e.text] = struct{}{}
	return 0
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
	return 0
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
	return 0
}

func (c *calendar) getEvents(w http.ResponseWriter, req *http.Request) {
	e := parseRequest(req)
	events := &getEventsModel{}

	var err int
	switch req.URL.Path[1:] {
	case "events_for_day":
		events, err = c.gatherEvents("day", e.date)
	case "events_for_week":
		events, err = c.gatherEvents("week", e.date)
	case "events_for_month":
		events, err = c.gatherEvents("month", e.date)
	default:
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	if err != 0 {
		w.WriteHeader(err)
		return
	}
	eventsJSON, errJSON := json.Marshal(events)
	if errJSON != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	w.Write(eventsJSON)
}

func (c *calendar) gatherEvents(howMuch, startDay string) (*getEventsModel, int) {
	model := &getEventsModel{Result: []DayEventsModel{}}
	var date, tillDay time.Time
	var err int

	switch howMuch {
	case "day":
		date, err = dayToTime(startDay)
		tillDay = date.AddDate(0, 0, 1)
	case "week":
		date, err = dayToTime(startDay)
		tillDay = date.AddDate(0, 0, 7)
	case "month":
		date, err = monthToTime(startDay)
		tillDay = date.AddDate(0, 1, 0)
	default:
		return nil, http.StatusServiceUnavailable
	}
	if err != 0 {
		return nil, http.StatusBadRequest
	}

	for ; date.Before(tillDay); date = date.AddDate(0, 0, 1) {
		thisDay := timeToDay(date)
		events := c.dayEvents(thisDay)
		if events != nil {
			model.Result = append(model.Result, *events)
		}
	}

	return model, 0
}

func (c *calendar) dayEvents(date string) *DayEventsModel {
	events, ok := c.days[date]
	if !ok || len(events) == 0 {
		return nil
	}
	allEvents := ""
	for text := range events {
		allEvents = allEvents + ", " + text
	}
	return &DayEventsModel{Date: date, Events: allEvents[2:]}
}

func dayToTime(date string) (time.Time, int) {
	dateTimed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, http.StatusBadRequest
	}
	return dateTimed, 0
}

func monthToTime(month string) (time.Time, int) {
	dateTimed, err := time.Parse("2006-01", month)
	if err != nil {
		return time.Time{}, http.StatusBadRequest
	}
	return dateTimed, 0
}

func timeToDay(day time.Time) string {
	date := day.Format("2006-01-02")
	return date
}
