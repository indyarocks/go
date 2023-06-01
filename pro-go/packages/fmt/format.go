package fmt

import (
	"fmt"
	"strconv"
)

func init() {
	fmt.Println("format.go init function invoked")
}

func ToCurrency(amount float64) string {
	return "$" + strconv.FormatFloat(amount, 'f', 2, 64)
}
