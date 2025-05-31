package concurrency

import (
	"context"
	"fmt"
	"sync"
)

func chanNotificator() {
	ch := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(i int) {
			<-ch
			fmt.Printf("Goroutine %d unlocked\n", i)
			wg.Done()
		}(i)
	}

	close(ch)
	wg.Wait()
}

func contextNotificator() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(i int) {
			<-ctx.Done()
			fmt.Printf("Goroutine %d unlocked\n", i)
			wg.Done()
		}(i)
	}

	cancel()
	wg.Wait()
}
