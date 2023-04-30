package main

import (
	"log"
	"net/http"
)

type calendar struct {
	days map[string]map[string]struct{}
}

type event struct {
	date    string
	text    string
	textOld string
}

func main() {
	c := &calendar{days: make(map[string]map[string]struct{})}
	config := getConfig()

	http.HandleFunc("/create_event", c.modifyEvent)
	http.HandleFunc("/update_event", c.modifyEvent)
	http.HandleFunc("/delete_event", c.modifyEvent)
	http.HandleFunc("/events_for_day", c.getEvents)
	http.HandleFunc("/events_for_week", c.getEvents)
	http.HandleFunc("/events_for_month", c.getEvents)
	log.Fatal(http.ListenAndServe(config.port, nil))
}
