package main

import (
	"fmt"
	"tempConvApp/tempconv"
)

func main() {
	fmt.Println(tempconv.CtoF(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.CtoF(tempconv.FreezingC))
	fmt.Println(tempconv.CtoF(tempconv.BoilingC))
}
