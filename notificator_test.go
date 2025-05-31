package concurrency

func ExampleChanNotificator() {
	chanNotificator()

	// Output:
	//Goroutine 2 unlocked
	//Goroutine 1 unlocked
	//Goroutine 0 unlocked
}

func ExampleContextNotificator() {
	contextNotificator()

	// Output:
	//Goroutine 2 unlocked
	//Goroutine 1 unlocked
	//Goroutine 0 unlocked
}
