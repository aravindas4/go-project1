package main

import (
	"fmt"
	"time"
)

type Log struct {
	Message string
	time.Time
}

func logger() {
	ts := time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC)
	log := Log{"Hello", ts}
	fmt.Printf("%#v\n", log)
}
