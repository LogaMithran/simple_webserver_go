package concurrency

import (
	"fmt"
	"sync"
)

var collectedDataFromTheApi = []int{}

func RaceCondition() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(3)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Routine 1 is writing")
		mut.Lock()
		collectedDataFromTheApi = append(collectedDataFromTheApi, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Routine 2 is writing")
		mut.Lock()
		collectedDataFromTheApi = append(collectedDataFromTheApi, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Routine 3 is writing")
		mut.Lock()
		collectedDataFromTheApi = append(collectedDataFromTheApi, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(collectedDataFromTheApi)
}
