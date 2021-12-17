package main

import (
	"example/user/hello/lower"
	"example/user/hello/maths"
	"example/user/hello/morestrings"
	"example/user/hello/upper"
	"fmt"
)

func main() {
	fmt.Println(morestrings.ReverseRunes(lower.LowerText("!oG ,olleH")))
	fmt.Println(morestrings.ReverseRunes(upper.UpperText("!oG ,olleH")))
	maths.CheckPrimeNumber(10)

}
