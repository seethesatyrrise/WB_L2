package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
	month   string
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

	http.HandleFunc("/create_event", c.createEvent)
	http.HandleFunc("/update_event", c.updateEvent)
	http.HandleFunc("/delete_event", c.deleteEvent)
	http.HandleFunc("/events_for_day", c.getEventsForDay)
	http.HandleFunc("/events_for_week", c.getEventsForWeek)
	http.HandleFunc("/events_for_month", c.getEventsForMonth)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseRequest(req *http.Request) *event {
	keys := req.URL.Query()
	e := &event{}

	if keys["date"] != nil {
		e.date = keys["date"][0]
	}
	if keys["month"] != nil {
		e.month = keys["month"][0]
	}
	if keys["textOld"] != nil {
		e.textOld = keys["textOld"][0]
	}
	if keys["text"] != nil {
		e.text = keys["text"][0]
	}
	return e
}

func (c *calendar) createEvent(w http.ResponseWriter, req *http.Request) {
	e := parseRequest(req)
	err := c.create(e)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, "error\n")
		return
	}
	io.WriteString(w, "created!\n")
}

func (c *calendar) create(e *event) error {
	if e.date == "" || e.text == "" {
		return errors.New("wrong args")
	}
	if c.days[e.date] == nil {
		c.days[e.date] = make(map[string]struct{})
	}
	c.days[e.date][e.text] = struct{}{}
	return nil
}

func (c *calendar) updateEvent(w http.ResponseWriter, req *http.Request) {
	e := parseRequest(req)
	err := c.update(e)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, "error\n")
		return
	}
	io.WriteString(w, "updated!\n")
}

func (c *calendar) update(e *event) error {
	if e.date == "" || e.text == "" || e.textOld == "" {
		return errors.New("wrong args")
	}
	_, ok := c.days[e.date][e.textOld]
	if !ok {
		return errors.New("can't find event " + e.textOld)
	}
	delete(c.days[e.date], e.textOld)
	c.days[e.date][e.text] = struct{}{}
	return nil
}

func (c *calendar) deleteEvent(w http.ResponseWriter, req *http.Request) {
	e := parseRequest(req)
	err := c.delete(e)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, "error\n")
		return
	}
	io.WriteString(w, "event deleted!\n")
}

func (c *calendar) delete(e *event) error {
	if e.date == "" || e.text == "" {
		return errors.New("wrong args")
	}
	_, ok := c.days[e.date][e.text]
	if !ok {
		return errors.New("can't find event " + e.text)
	}
	delete(c.days[e.date], e.text)
	return nil
}

func (c *calendar) getEventsForDay(w http.ResponseWriter, req *http.Request) {
	e := parseRequest(req)
	events, err := c.dayEvents(e.date)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, "error\n")
		return
	}
	res := getEventsModel{Result: []DayEventsModel{*events}}
	eventsJSON, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, "error\n")
		return
	}
	w.Write(eventsJSON)
}

func (c *calendar) dayEvents(date string) (*DayEventsModel, error) {
	events, ok := c.days[date]
	if !ok || len(events) == 0 {
		return nil, errors.New("can't find events for day " + date)
	}
	res := ""
	for text := range events {
		res = res + ", " + text
	}
	return &DayEventsModel{Date: date, Events: res[2:]}, nil
}

func (c *calendar) getEventsForWeek(w http.ResponseWriter, req *http.Request) {
	e := parseRequest(req)
	events, err := c.weekEvents(e.date)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, "error\n")
		return
	}
	eventsJSON, err := json.Marshal(events)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, "error\n")
		return
	}
	w.Write(eventsJSON)
}

func (c *calendar) weekEvents(date string) (*getEventsModel, error) {
	model := &getEventsModel{Result: []DayEventsModel{}}
	day, err := dayToTime(date)
	if err != nil {
		return nil, err
	}
	nextWeek := day.AddDate(0, 0, 7)

	for ; day.Before(nextWeek); day = day.AddDate(0, 0, 1) {
		date = timeToDay(day)
		events, ok := c.days[date]
		if !ok || len(events) == 0 {
			continue
		}
		res := ""
		for text := range events {
			res = res + ", " + text
		}
		model.Result = append(model.Result, DayEventsModel{Date: date, Events: res[2:]})

	}

	return model, nil
}

func dayToTime(date string) (time.Time, error) {
	dateTimed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}
	return dateTimed, nil
}

func monthToTime(month string) (time.Time, error) {
	dateTimed, err := time.Parse("2006-01", month)
	if err != nil {
		return time.Time{}, err
	}
	return dateTimed, nil
}

func timeToDay(day time.Time) string {
	date := day.Format("2006-01-02")
	return date
}

func (c *calendar) getEventsForMonth(w http.ResponseWriter, req *http.Request) {
	e := parseRequest(req)
	events, err := c.monthEvents(e.month)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, "error\n")
		return
	}
	eventsJSON, err := json.Marshal(events)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, "error\n")
		return
	}
	w.Write(eventsJSON)
}

func (c *calendar) monthEvents(month string) (*getEventsModel, error) {
	model := &getEventsModel{Result: []DayEventsModel{}}
	day, err := monthToTime(month)
	if err != nil {
		return nil, err
	}
	nextMonth := day.AddDate(0, 1, 0)

	for ; day.Before(nextMonth); day = day.AddDate(0, 0, 1) {
		date := timeToDay(day)
		events, ok := c.days[date]
		if !ok || len(events) == 0 {
			continue
		}
		res := ""
		for text := range events {
			res = res + ", " + text
		}
		model.Result = append(model.Result, DayEventsModel{Date: date, Events: res[2:]})
	}

	return model, nil
}
