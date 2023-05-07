package rest

import (
	"WB-L2/develop/dev11/internal/models"
	"WB-L2/develop/dev11/internal/utils"
	"errors"
	"net/http"
	"time"
)

func (c *Calendar) EventsForDay(w http.ResponseWriter, req *http.Request) {
	c.GetEvents(w, req, "day", &eventNeeds{date: true})
}

func (c *Calendar) EventsForWeek(w http.ResponseWriter, req *http.Request) {
	c.GetEvents(w, req, "week", &eventNeeds{date: true})
}

func (c *Calendar) EventsForMonth(w http.ResponseWriter, req *http.Request) {
	c.GetEvents(w, req, "month", &eventNeeds{month: true})
}

func (c *Calendar) GetEvents(w http.ResponseWriter, req *http.Request, howMuch string, needs *eventNeeds) {
	keys := req.URL.Query()
	e, err := ParseRequest(keys, needs)
	if err != nil {
		utils.PublishError(w, http.StatusBadRequest, err)
		return
	}

	events, status, err := c.gatherEvents(howMuch, e.date)
	if err != nil {
		utils.PublishError(w, status, err)
		return
	}

	utils.PublishData(w, status, events)
}

func (c *Calendar) gatherEvents(howMuch, startDay string) (*models.GetEventsModel, int, error) {
	model := &models.GetEventsModel{Result: []models.DayEventsModel{}}
	var date, tillDay time.Time
	var status int
	var err error

	switch howMuch {
	case "day":
		date, status, err = utils.DayToTime(startDay)
		tillDay = date.AddDate(0, 0, 1)
	case "week":
		date, status, err = utils.DayToTime(startDay)
		tillDay = date.AddDate(0, 0, 7)
	case "month":
		date, status, err = utils.MonthToTime(startDay)
		tillDay = date.AddDate(0, 1, 0)
	default:
		return nil, http.StatusServiceUnavailable, errors.New("something weird")
	}
	if err != nil {
		return nil, status, err
	}

	for ; date.Before(tillDay); date = date.AddDate(0, 0, 1) {
		thisDay := utils.TimeToDay(date)
		events := c.dayEvents(thisDay)
		if events != nil {
			model.Result = append(model.Result, *events)
		}
	}

	return model, http.StatusOK, nil
}

func (c *Calendar) dayEvents(date string) *models.DayEventsModel {
	events, ok := c.days[date]
	if !ok || len(events) == 0 {
		return nil
	}
	allEvents := ""
	for text := range events {
		allEvents = allEvents + ", " + text
	}
	return &models.DayEventsModel{Date: date, Events: allEvents[2:]}
}
