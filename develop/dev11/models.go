package main

type getEventsModel struct {
	Result []DayEventsModel `json:"result"`
}

type DayEventsModel struct {
	Date   string `json:"date"`
	Events string `json:"events"`
}

type errorModel struct {
	Err string `json:"error"`
}

type responseModel struct {
	Resp string `json:"result"`
}
