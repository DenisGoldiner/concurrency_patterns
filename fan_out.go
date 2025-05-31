package concurrency

func fanOut(in <-chan int) []chan int {
	var out []chan int

	for i := 0; i < 3; i++ {
		out = append(out, make(chan int, 1))
	}

	go func() {
		for v := range in {
			switch {
			case v%3 == 0:
				out[2] <- v
			case v%2 == 0:
				out[1] <- v
			default:
				out[0] <- v
			}
		}

		for _, ch := range out {
			close(ch)
		}
	}()

	return out
}
