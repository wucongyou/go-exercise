package executor

import (
	"fmt"
	"sync"
	"time"
)

// Callable callable interface.
type Callable func() (interface{}, error)

const (
	duration time.Duration = time.Millisecond * 100
)

// Start start a example.
func Start() {
	ch := make(chan Callable, 10240)
	var wg sync.WaitGroup
	endCh := make(chan int, 1)
	wg.Add(2)
	go func() {
		defer wg.Done()
		produceproc(endCh, ch)
	}()
	go func() {
		defer wg.Done()
		consumeproc(endCh, ch)
	}()
	wg.Wait()
}

func consumeproc(endCh chan int, ch chan Callable) {
	var (
		callErr error
		res     interface{}
	)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic(%v) detected, restarting\n", err)
			go consumeproc(endCh, ch)
		}
	}()
	for {
		select {
		case <-endCh:
			fmt.Println("consumer got a end signal, exit")
			return
		default:
		}
		time.Sleep(duration)
		callable := <-ch
		if res, callErr = callable(); callErr != nil {
			fmt.Printf("error(%v)\n", callErr)
		}
		fmt.Printf("res: %v\n", res)
	}
}

func produceproc(endCh chan int, ch chan Callable) {
	count := 0
	for i := 0; i < 10; i++ {
		time.Sleep(duration)
		count++
		ch <- func() (interface{}, error) {
			return func() (res int, err error) {
				if count%3 == 0 {
					panic(fmt.Sprintf("count: %d", count))
				}
				if count%5 == 0 {
					err = fmt.Errorf("count: %d", count)
					return
				}
				res = count
				return
			}()
		}
	}
	endCh <- 1
}
