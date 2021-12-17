package maths

import (
	"fmt"
	"math"
)

func CheckPrimeNumber(num int) {
	if num < 2 {
		fmt.Println("Number must be greater than 2.")
		return
	}
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			fmt.Println("Non Prime Number")
			return
		}
	}
	fmt.Println("Prime Number")
	return
}
