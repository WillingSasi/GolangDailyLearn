package main

import (
	"fmt"
	"math"
)

type dualError struct {
	Num     float64
	problem string
}

func (e dualError) Error() string {
	return fmt.Sprintf("Wrong!!! Because \"%f\" is a negatice number", e.Num)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// return -1, errors.New("math: square root of negative number")
		return -1, dualError{Num: f}
	}

	return math.Sqrt(f), nil
}

func main() {
	result, err := Sqrt(-13)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
