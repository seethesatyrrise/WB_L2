package main

import "time"

func validateDay(day string) bool {
	_, err := time.Parse("2006-01-02", day)
	if err != nil {
		return false
	}
	return true
}

func validateMonth(month string) bool {
	_, err := time.Parse("2006-01", month)
	if err != nil {
		return false
	}
	return true
}
