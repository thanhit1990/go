package upper

import (
	"fmt"
	"strings"
)

func UpperText(s string) string {
	fmt.Println(s)
	res := strings.ToUpper(s)
	fmt.Println(res)
	return res
}
