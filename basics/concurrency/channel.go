package concurrency

import (
	"fmt"
	"sync"
)

func TestChannel() {
	/**
	we can pass the buffer value for as the second param
	Meaning that only 2 write are gonna happen
	*/

	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// <- represents the out channel, ie) it's going to emit the value
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	// here <- represents it's getting the value
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 5
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
