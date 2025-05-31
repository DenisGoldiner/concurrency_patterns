package concurrency

import (
	"fmt"
	"sync"
)

func workers(in <-chan int) chan int {
	const n = 3
	outCh := make(chan int, n)
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			for val := range in {
				outCh <- val
				fmt.Printf("worker %d processed value %d\n", i, val)
			}
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}
