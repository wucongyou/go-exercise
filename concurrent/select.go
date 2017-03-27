package concurrent

import (
	"fmt"
	"sync"
)

func Select() {
	var (
		wg   sync.WaitGroup
		a, b = make(chan int), make(chan int)
	)
	go func() {
		defer wg.Done()
		for {
			var (
				name string
				ok   bool
				x    int
			)
			select {
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}
			if !ok {
				return
			}
			fmt.Printf("%d from %s\n", x, name)
		}
	}()
	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)
		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			case b <- i * 10:
			}
		}
	}()
	wg.Wait()

}
