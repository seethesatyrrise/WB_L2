package server

import (
	"WB-L2/develop/dev11/internal/rest"
	"WB-L2/develop/dev11/internal/utils"
	"context"
	"fmt"
	"net/http"
	"time"
)

func NewServer(c *rest.Calendar, config *PortConfig) *http.Server {

	mux := http.NewServeMux()
	mux.Handle("/create_event", utils.Logging(http.HandlerFunc(c.CreateEvent)))
	mux.Handle("/update_event", utils.Logging(http.HandlerFunc(c.UpdateEvent)))
	mux.Handle("/delete_event", utils.Logging(http.HandlerFunc(c.DeleteEvent)))
	mux.Handle("/events_for_day", utils.Logging(http.HandlerFunc(c.EventsForDay)))
	mux.Handle("/events_for_week", utils.Logging(http.HandlerFunc(c.EventsForWeek)))
	mux.Handle("/events_for_month", utils.Logging(http.HandlerFunc(c.EventsForMonth)))

	s := &http.Server{
		Addr:           config.Port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s
}

func Start(s *http.Server) error {
	go func() {
		if err := s.ListenAndServe(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	fmt.Println("server started")
	return nil
}

func Shutdown(ctx context.Context, s *http.Server) error {
	if s != nil {
		return s.Shutdown(ctx)
	}

	fmt.Println("server stopped")
	return nil
}
