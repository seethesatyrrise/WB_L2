package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

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
