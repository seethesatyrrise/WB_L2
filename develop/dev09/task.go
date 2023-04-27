package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var defaultNameCounter int64 = 0

type command struct {
	com  string
	url  string
	file string
}

type file struct {
	names []string
	urls  []string
}

func main() {
	fmt.Println("input url (and optionally an output file name) or name of the file with the list of URLs\n" +
		"examples: \"url https://www.google.com/ google\" \"file urls\"")
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		c, err := parseInput(input)
		if err != nil {
			fmt.Println(err)
		}
		c.wget()
	}
}

func parseInput(com string) (*command, error) {
	c := &command{}
	w := strings.Fields(com)
	switch w[0] {
	case "url":
		if len(w) <= 1 {
			return nil, errors.New("not enough arguments")
		}
		c.com = "url"
		c.url = w[1]
		if len(w) > 2 {
			c.file = w[2]
		}
		if len(w) > 3 {
			fmt.Println("too many names, I will use first of them")
		}
	case "file":
		if len(w) <= 1 {
			return nil, errors.New("not enough arguments")
		}
		c.com = "file"
		c.file = w[1]
		if len(w) > 2 {
			fmt.Println("too many names, I will use first of them")
		}
	default:
		fmt.Println("undefined command")
	}

	return c, nil
}

func (c *command) wget() error {
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
		f.getManyUrls()
	}
	return nil
}

func parseFile(name string) (*file, error) {
	newFile, err := os.Open("develop\\dev09\\" + name)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer newFile.Close()

	f := &file{}

	scanner := bufio.NewScanner(newFile)
	for scanner.Scan() {
		line := scanner.Text()
		w := strings.Fields(line)
		if len(w) < 1 {
			continue
		}
		f.urls = append(f.urls, w[0])
		if len(w) > 1 {
			f.names = append(f.names, w[1])
		} else {
			f.names = append(f.names, "")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return f, nil
}

func (f *file) getManyUrls() {
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

func saveInFile(resp *http.Response, url, name string) {
	if name == "" {
		name = getDefaultName()
	}
	file, err := os.Create("develop\\dev09\\" + name)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("body of the " + url + " was saved in file " + name)
	return
}

func getDefaultName() string {
	counter := strconv.Itoa(int(defaultNameCounter))
	atomic.AddInt64(&defaultNameCounter, 1)
	return "url" + counter
}
