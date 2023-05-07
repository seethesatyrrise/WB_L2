package utils

import (
	"net/http"
	"time"
)

func DayToTime(date string) (time.Time, int, error) {
	dateTimed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, http.StatusBadRequest, err
	}
	return dateTimed, http.StatusOK, nil
}

func MonthToTime(month string) (time.Time, int, error) {
	dateTimed, err := time.Parse("2006-01", month)
	if err != nil {
		return time.Time{}, http.StatusBadRequest, err
	}
	return dateTimed, http.StatusOK, nil
}

func TimeToDay(day time.Time) string {
	date := day.Format("2006-01-02")
	return date
}
