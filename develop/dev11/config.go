package main

import "os"

type portConfig struct {
	port string
}

func getConfig() *portConfig {
	port := os.Getenv("port")
	if port == "" {
		port = ":8080"
	}
	return &portConfig{port: port}
}
