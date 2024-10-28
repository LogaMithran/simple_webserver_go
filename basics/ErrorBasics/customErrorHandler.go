package ErrorBasics

import (
	"fmt"
	"math"
)

type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("Error: Cannot calculate SQRT for negative number %g", float64(f))
}

func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, NegativeSqrtError(a)
	}
	return math.Sqrt(a), nil
}

func CalculateError() {
	a := 16
	if result, err := sqrt(float64(a)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt of ", a, " is ", result)
	}

	a = -16

	if result, err := sqrt(float64(a)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt of ", a, " is ", result)
	}
}
