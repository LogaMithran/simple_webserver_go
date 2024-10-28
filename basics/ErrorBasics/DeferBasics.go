package ErrorBasics

import "fmt"

func DeferBasics() {
	fmt.Println("\nDefer basics\n")

	/** Defer always works on last in first out concept */
	for a := 0; a <= 5; a++ {
		defer fmt.Println(a)
	}
	fmt.Println()
}
