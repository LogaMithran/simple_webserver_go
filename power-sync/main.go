package main

import (
	"fmt"
	"power-sync/people"
)

func main() {
	fmt.Println("Started syncing people from SWAPI...")
	people.StartSyncingPeople()
}
