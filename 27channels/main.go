package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang.")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	// READ ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(myCh)
		val, IsChannelOpen := <-myCh

		fmt.Println(IsChannelOpen)
		if IsChannelOpen {
			fmt.Println(val)
		} else {
			fmt.Println("Channel is closed.")
			fmt.Println("val")
		}

		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// SEND ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 0
		// myCh <- 5
		// myCh <- 6
		// close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
