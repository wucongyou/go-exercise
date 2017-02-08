package main

import (
	"fmt"
	"time"
)

type Callable func() (interface{}, error)

const (
	duration time.Duration = 1000
)

func main() {
	ch := make(chan Callable, 10240)
	go producer(ch)
	go consumer(ch)
	time.Sleep(100 * time.Second)
}

func consumer(ch chan Callable) {
	for {
		time.Sleep(duration)
		callable := <-ch
		fmt.Printf("consume start: ")
		count, _ := callable()
		fmt.Printf("consumed: %d\n", count)
	}
}

func producer(ch chan Callable) {
	var count int = 0
	for {
		time.Sleep(duration)
		count += 1
		fmt.Printf("produce start\n")
		ch <- func() (interface{}, error) {
			newCount, err := func() (int, error) {
				fmt.Printf("%d\n", count)
				return count, nil
			}()
			return newCount, err
		}
	}
}
