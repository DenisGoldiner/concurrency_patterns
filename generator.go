package concurrency

import (
	"fmt"
	"sync"
)

func fabricsOut() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}

		close(ch)
	}()

	return ch
}

func fabricsIn(wg *sync.WaitGroup) chan<- int {
	ch := make(chan int)

	wg.Add(1)
	go func() {
		for v := range ch {
			fmt.Printf("Recieved value %d\n", v)
		}

		wg.Done()
	}()

	return ch
}
