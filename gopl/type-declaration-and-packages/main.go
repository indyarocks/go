package main

import (
	"fmt"
	"os"
	"strconv"
	"tempConvApp/tempconv"
)

func main() {
	fmt.Println(tempconv.CtoF(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.CtoF(tempconv.FreezingC))
	fmt.Println(tempconv.CtoF(tempconv.BoilingC))

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celcius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FtoC(f), c, tempconv.CtoF(c))
	}

	medals := []string{"gold", "silver", "bronze"}
	fmt.Printf("%T\n", medals)
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
}
