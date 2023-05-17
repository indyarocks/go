package main

import (
	"fmt"
	"sort"
)

func main() {
	names := [4]string{"Chandan", "Tithi", "Darsh", "Sheenu"}

	secondName := names[1]
	fmt.Println(secondName)
	sort.Strings(names[:])
	fmt.Println(names[:])
	fmt.Println(secondName)
	fmt.Println("Now using pointers...")
	names1 := [4]string{"Chandan", "Tithi", "Darsh", "Sheenu"}
	secondPosition := &names1[1]
	fmt.Println(*secondPosition)
	sort.Strings(names1[:])
	fmt.Println(names[:])
	fmt.Println(*secondPosition)
}
