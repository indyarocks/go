package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "Use if you want to omit trailing newline")
var sep = flag.String("s", " ", "Separator for the arguments")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
