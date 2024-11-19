package concurrency

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Start() {
	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()
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
	i := 1
	for api := range starWarsApi {
		wg.Add(i)
		i++
		go genericFuncForCallingApi(starWarsApi[api])
	}

	wg.Wait()
}

func genericFuncForCallingApi(url string) int {
	defer wg.Done()
	if res, err := http.Get(url); err != nil {
		fmt.Println("Some error in calling the API", url)
	} else {
		fmt.Println("Response from the API", res.StatusCode, res.Body)
	}

	return 1
}
