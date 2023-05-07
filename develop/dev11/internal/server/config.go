package server

import "os"

type PortConfig struct {
	Port string
}

func GetConfig() *PortConfig {
	port := os.Getenv("port")
	if port == "" {
		port = ":8080"
	}
	return &PortConfig{Port: port}
}
