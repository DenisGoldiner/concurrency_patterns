package concurrency

import "fmt"

func ExampleFanOut() {
	in := make(chan int)
	outs := fanOut(in)

	go func() {
		in <- 1
		in <- 2
		in <- 3

		close(in)
	}()

	for i, ch := range outs {
		for val := range ch {
			fmt.Printf("Value from the chan #%d = %d\n", i, val)
		}
	}

	// Output:
	//Value from the chan #0 = 1
	//Value from the chan #1 = 2
	//Value from the chan #2 = 3
}
