package concurrency

import (
	"fmt"
	"sync"
)

func fanIn(ins []chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int, len(ins))

	wg.Add(len(ins))
	for i, in := range ins {
		go func(i int, in <-chan int) {
			for val := range in {
				out <- val
				fmt.Printf("Value from the chan #%d = %d was merged\n", i, val)
			}
			wg.Done()
		}(i, in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
