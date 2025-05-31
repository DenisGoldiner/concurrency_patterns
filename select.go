package concurrency

func syncSelector(ch1, ch2, ch3 <-chan int) int {
	select {
	case v := <-ch1:
		return v
	case v := <-ch2:
		return v
	case v := <-ch3:
		return v
	}
}

func asyncSelector(ch1, ch2, ch3 <-chan int) int {
	select {
	case v := <-ch1:
		return v
	case v := <-ch2:
		return v
	case v := <-ch3:
		return v
	default:
		return 0
	}
}
