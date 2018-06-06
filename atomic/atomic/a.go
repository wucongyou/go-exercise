package main

import "sync/atomic"

func main() {
	var a atomic.Value
	endpoints := make([]string, 0)
	a.Store(endpoints)
	endpoints = a.Load().([]string)
}
