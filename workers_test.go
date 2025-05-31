package concurrency

func ExampleWorkers() {
	in := make(chan int)

	out := workers(in)

	go func() {
		for i := 0; i < 3; i++ {
			in <- i
		}

		close(in)
	}()

	for range out {
	}

	// Output:
	//worker 1 processed value 0
	//worker 0 processed value 2
	//worker 2 processed value 1
}
