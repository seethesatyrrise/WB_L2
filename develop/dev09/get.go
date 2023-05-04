package main

import (
	"fmt"
	"net/http"
	"sync"
)

type command struct {
	com  string
	url  string
	file string
}

type file struct {
	names []string
	urls  []string
}

func (c *command) run() error {
	switch c.com {
	case "url":
		resp, err := getUrl(c.url)
		if err != nil {
			return err
		}
		saveInFile(resp, c.url, c.file)
	case "file":
		f, err := parseFile(c.file)
		if err != nil {
			return err
		}
		f.wget()
	}
	return nil
}

func (f *file) wget() {
	wg := sync.WaitGroup{}

	for i, url := range f.urls {
		wg.Add(1)
		go func(i int, url string) {
			resp, err := getUrl(url)
			if err != nil {
				fmt.Println(err)
			} else {
				saveInFile(resp, url, f.names[i])
			}
			wg.Done()
		}(i, url)
	}

	wg.Wait()
}

func getUrl(url string) (*http.Response, error) {
	return http.Get(url)
}
