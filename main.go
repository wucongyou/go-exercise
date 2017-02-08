package main

import (
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
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("error(%v) detected, restarting\n", err)
			go consumer(ch)
		}
	}()
	for {
		time.Sleep(duration)
		callable := <-ch
		count, _ := callable()
		fmt.Printf("%d\n", count)
	}
}

func producer(ch chan Callable) {
	var count int = 0
	for {
		time.Sleep(duration)
		count += 1
		ch <- func() (interface{}, error) {
			newCount, err := func() (int, error) {
				if count%5 == 0 {
					panic(fmt.Sprintf("count: %d, panic", count))
				}
				return count, nil
			}()
			return newCount, err
		}
	}
}
