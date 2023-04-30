package main

import (
	"errors"
	"net/http"
	"time"
)

func (c *calendar) getEvents(w http.ResponseWriter, req *http.Request) {
	e, err := parseRequest(req)
	if err != nil {
		publishError(w, http.StatusBadRequest, err)
		return
	}

	events := &getEventsModel{}
	var status int

	switch req.URL.Path[1:] {
	case "events_for_day":
		events, status, err = c.gatherEvents("day", e.date)
	case "events_for_week":
		events, status, err = c.gatherEvents("week", e.date)
	case "events_for_month":
		events, status, err = c.gatherEvents("month", e.date)
	default:
		publishError(w, http.StatusServiceUnavailable, errors.New("???"))
		return
	}
	if err != nil {
		publishError(w, status, err)
		return
	}

	publishData(w, status, events)
}

func (c *calendar) gatherEvents(howMuch, startDay string) (*getEventsModel, int, error) {
	model := &getEventsModel{Result: []DayEventsModel{}}
	var date, tillDay time.Time
	var status int
	var err error

	switch howMuch {
	case "day":
		date, status, err = dayToTime(startDay)
		tillDay = date.AddDate(0, 0, 1)
	case "week":
		date, status, err = dayToTime(startDay)
		tillDay = date.AddDate(0, 0, 7)
	case "month":
		date, status, err = monthToTime(startDay)
		tillDay = date.AddDate(0, 1, 0)
	default:
		return nil, http.StatusServiceUnavailable, errors.New("can't ?")
	}
	if err != nil {
		return nil, status, err
	}

	for ; date.Before(tillDay); date = date.AddDate(0, 0, 1) {
		thisDay := timeToDay(date)
		events := c.dayEvents(thisDay)
		if events != nil {
			model.Result = append(model.Result, *events)
		}
	}

	return model, http.StatusOK, nil
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
