package main

import (
	"strconv"
	"sync/atomic"
)

var defaultNameCounter int64 = 0

func getDefaultName() string {
	counter := strconv.Itoa(int(defaultNameCounter))
	atomic.AddInt64(&defaultNameCounter, 1)

	return "url" + counter
}
