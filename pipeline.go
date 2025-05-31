package concurrency

import (
	"fmt"
)

func pipeline() {
	leftmost := make(chan int)
	right, left := leftmost, leftmost

	for i := 0; i < 3; i++ {
		right = make(chan int)

		go func(left, right chan int) {
			val := <-right
			fmt.Printf("Link %d processed\n", val)
			left <- val + 1
		}(left, right)

		left = right
	}

	go func(c chan int) { c <- 0 }(right)

	<-leftmost
}
