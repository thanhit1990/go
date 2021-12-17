package main

import (
	"fmt"

	"github.com/thanhit1990/go/lower"
	"github.com/thanhit1990/go/maths"
	"github.com/thanhit1990/go/morestrings"
	"github.com/thanhit1990/go/upper"
)

func main() {
	fmt.Println(morestrings.ReverseRunes(lower.LowerText("!oG ,olleH")))
	fmt.Println(morestrings.ReverseRunes(upper.UpperText("!oG ,olleH")))
	maths.CheckPrimeNumber(10)

}
