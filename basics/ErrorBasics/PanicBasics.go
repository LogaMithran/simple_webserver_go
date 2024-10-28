package ErrorBasics

import "fmt"

func PanicBasics() {
	fmt.Println("Started")
	call_func_a()
	fmt.Println("Ended")
}

func call_func_a() {
	fmt.Println("In function A")
	defer func() {
		// If recover is not added then the function will break all together
		if revived := recover(); revived != nil {
			fmt.Println("Recovered from function A")
		}
	}()

	call_func_b(0)
}

func call_func_b(a int) {
	fmt.Println("\nIn function B")
	if a > 3 {
		panic(fmt.Sprintf("Panicking from %v", a))
	}
	defer fmt.Printf("Defering %v", a)

	fmt.Printf("Printing in func B %v", a)
	call_func_b(a + 1)
}
