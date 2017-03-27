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
			time.Sleep(time.Second)
			fmt.Printf("goroutine %d done\n", id)
		}(i)
	}
	fmt.Println("main waiting")
	wg.Wait()
	fmt.Println("main done")
}
