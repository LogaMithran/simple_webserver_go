package people

import (
	"fmt"
	"strconv"
)

func StartSyncingPeople() {
	var syncCount = 1

	for i := range 100 {
		if response := Call("https://swapi.dev/api/people/" + strconv.Itoa(i)); response.StatusCode == 404 {
			fmt.Printf("Not found %d", i)
		} else {
			fmt.Println(response.StatusCode)
		}
		syncCount += 1
	}
	fmt.Printf("Total peoples synced %d", syncCount-1)
}
