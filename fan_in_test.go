package concurrency

func ExampleFanIn() {
	var ins []chan int

	for i := 0; i < 3; i++ {
		in := make(chan int, 1)
		ins = append(ins, in)
	}

	out := fanIn(ins)

	go func() {
		for i, in := range ins {
			in <- i
			close(in)
		}
	}()

	for range out {
		// do nothing
	}

	// Output:
	//Value from the chan #2 = 2 was merged
	//Value from the chan #1 = 1 was merged
	//Value from the chan #0 = 0 was merged
}
