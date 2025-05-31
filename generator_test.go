package concurrency

import (
	"fmt"
	"sync"
)

func ExampleFabricsOut() {
	ch := fabricsOut()

	for v := range ch {
		fmt.Printf("Recieved value %d\n", v)
	}

	// Output:
	//Recieved value 0
	//Recieved value 1
	//Recieved value 2
}

func ExampleFabricsIn() {
	wg := sync.WaitGroup{}
	ch := fabricsIn(&wg)

	for i := 0; i < 3; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()

	// Output:
	//Recieved value 0
	//Recieved value 1
	//Recieved value 2
}
