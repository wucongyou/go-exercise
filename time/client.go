package main

import (
	"fmt"
	"time"
)

const (
	_timeFormat = "2006-01-02 15:04:05"
)

func main() {
	s := "2017-12-28 20:19:05"

	ti, err := time.ParseInLocation(_timeFormat, s, time.Now().Location())
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
	ts := ti.Unix()
	fmt.Printf("ti: %v, ts: %d", ti, ts)
}
