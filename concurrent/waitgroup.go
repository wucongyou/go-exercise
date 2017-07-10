package concurrent

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			doSomething()
			fmt.Printf("work goroutine %d done\n", id)
		}(i)
	}
	fmt.Println("main goroutine waiting")
	wg.Wait()
	fmt.Println("main goroutine done")
}

func doSomething() {
	time.Sleep(time.Second)
}
