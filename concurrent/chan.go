package concurrent

import (
	"fmt"
	"sync"
	"time"
)

func Chan() {
	var wg sync.WaitGroup
	ready := make(chan []struct{})
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("%d : ready", id)
			<-ready
			fmt.Printf("%d : runnning", id)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Ready? Go!")
	close(ready)
	wg.Wait()
}
