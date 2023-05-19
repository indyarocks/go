package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	intMax := math.MaxInt64
	floatMax := math.MaxFloat64
	fmt.Println(intMax)
	fmt.Println(intMax * 2)
	fmt.Println(floatMax)
	fmt.Println(floatMax * 2)
	fmt.Println(math.IsInf((floatMax * 2), 0))

	first := 100
	const second = 200.00
	fmt.Println(first < second)
	fmt.Println()

	falseStringSlice := []string{"0", "true", "1", "TRUE", "meNotTrue"}
	for _, str := range falseStringSlice {
		if boolVal, boolErr := strconv.ParseBool(str); boolErr == nil {
			fmt.Println("Parsed value of ", str, " is ", boolVal)
		} else {
			fmt.Println("Parsing Error for ", str, " And error is ", boolErr)
		}
	}

	binaryValStringSlice := []string{"100", "11", "11111", "101010", "11111", "1", "10", "11111111111", "99"}
	for _, binaryValString := range binaryValStringSlice {
		int1, int1err := strconv.ParseInt(binaryValString, 2, 8)

		if int1err == nil {
			smallInt := int8(int1)
			fmt.Println("Parsed value to binaryValString", binaryValString, "is", smallInt)
		} else {
			fmt.Println("Can't parse", int1, int1err)
		}
	}
}
