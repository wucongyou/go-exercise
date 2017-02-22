package main

import (
	"errors"
	"fmt"
	"time"
)

type Callable func() (interface{}, error)

const (
	duration time.Duration = 1 * time.Second
)

func main() {
	ch := make(chan Callable, 10240)
	go producer(ch)
	go consumer(ch)
	time.Sleep(100 * time.Second)
}

func consumer(ch chan Callable) {
	var (
		callErr error
		res     interface{}
	)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic(%v) detected, restarting\n", err)
			go consumer(ch)
		}
	}()
	for {
		time.Sleep(duration)
		callable := <-ch
		if res, callErr = callable(); callErr != nil {
			fmt.Printf("error(%v)\n", callErr)
		}
		fmt.Printf("res: %v\n", res)
	}
}

func producer(ch chan Callable) {
	var count int = 0
	for {
		time.Sleep(duration)
		count += 1
		ch <- func() (interface{}, error) {
			return func() (res int, err error) {
				if count%3 == 0 {
					panic(fmt.Sprintf("count: %d", res))
				}
				if count%5 == 0 {
					err = errors.New(fmt.Sprintf("count: %d", res))
					return
				}
				res = count
				return
			}()
		}
	}
}
