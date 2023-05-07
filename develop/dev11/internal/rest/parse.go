package rest

import (
	"WB-L2/develop/dev11/internal/utils"
	"errors"
	"net/url"
)

func ParseRequest(keys url.Values, needs *eventNeeds) (*event, error) {
	e := &event{}
	var err error

	if needs.date {
		e.date, err = getDate(keys)
		if err != nil {
			return nil, err
		}
	}

	if needs.text {
		e.text, err = getText(keys, "text")
		if err != nil {
			return nil, err
		}
	}

	if needs.month {
		e.date, err = getMonth(keys)
		if err != nil {
			return nil, err
		}
	}

	if needs.textOld {
		e.textOld, err = getText(keys, "textOld")
		if err != nil {
			return nil, err
		}
	}

	return e, nil
}

func getDate(keys url.Values) (string, error) {
	var date string
	if keys["date"] != nil {
		date = keys["date"][0]
		if !utils.ValidateDay(date) {
			return "", errors.New("wrong date string")
		}
	}
	return date, nil
}

func getMonth(keys url.Values) (string, error) {
	var date string
	if keys["month"] != nil {
		date = keys["month"][0]
		if !utils.ValidateMonth(date) {
			return "", errors.New("wrong month string")
		}
	}
	return date, nil
}

func getText(keys url.Values, eventText string) (string, error) {
	var text string
	if keys[eventText] != nil {
		text = keys[eventText][0]
		if text == "" {
			return "", errors.New("empty event " + eventText + " string")
		}
	}
	return text, nil
}
