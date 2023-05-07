package rest

type Calendar struct {
	days map[string]map[string]struct{}
}

type event struct {
	date    string
	text    string
	textOld string
}

type eventNeeds struct {
	date    bool
	month   bool
	text    bool
	textOld bool
}

func NewCalendar() *Calendar {
	return &Calendar{days: make(map[string]map[string]struct{})}
}
