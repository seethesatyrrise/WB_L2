package models

type GetEventsModel struct {
	Result []DayEventsModel `json:"result"`
}

type DayEventsModel struct {
	Date   string `json:"date"`
	Events string `json:"events"`
}

type ErrorModel struct {
	Err string `json:"error"`
}

type ResponseModel struct {
	Resp string `json:"result"`
}
