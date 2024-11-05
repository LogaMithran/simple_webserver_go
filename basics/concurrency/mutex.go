package concurrency

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var waitg sync.WaitGroup
var mut sync.Mutex

var webStatus = []int{}

func StartMutex() {
	starWarsApi := []string{
		"https://swapi.dev/api/people/1",
		"https://swapi.dev/api/people/2",
		"https://swapi.dev/api/people/3",
		"https://swapi.dev/api/people/4",
		"https://swapi.dev/api/people/5",
		"https://swapi.dev/api/people/6",
	}
	t := time.Now()
	defer func() {
		fmt.Println(time.Since(t))
	}()
	for api := range starWarsApi {
		go callApiAndWriteValue(starWarsApi[api])
		waitg.Add(1)
	}
	waitg.Wait()
	fmt.Println(webStatus)
}

func callApiAndWriteValue(url string) int {
	defer waitg.Done()
	if res, err := http.Get(url); err != nil {
		fmt.Println("Some error in calling the API", url)
	} else {
		fmt.Println("Response from the API", res.StatusCode, res.Body)
		mut.Lock()
		webStatus = append(webStatus, res.StatusCode)
		mut.Unlock()
	}

	return 1
}
