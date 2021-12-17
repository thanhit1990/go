package lower

import (
	"fmt"
	"strings"
)

func LowerText(s string) string {
	fmt.Println(s)
	res := strings.ToLower(s)
	fmt.Println(res)
	return res
}
