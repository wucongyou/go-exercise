package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan []int, 10)
	var arr []int
	var mu sync.Mutex
	go func() {
		for {
			mu.Lock()
			res := <-ch
			fmt.Printf("arr: %v, res: %v\n", arr, res)
			mu.Unlock()
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < 10; i++ {
		mu.Lock()
		arr = []int{i}
		ch <- arr
		arr = []int{11}
		mu.Unlock()
	}
	time.Sleep(time.Second * 10)
}
