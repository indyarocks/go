package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Chandan"
	names[1] = "Tithi"
	names[2] = "Darsh"

	for i, name := range names {
		fmt.Printf("names[%d]=%s\n", i, name)
	}

	var surNames [3]string = names
	fmt.Println(surNames)

	var unKnowns = [...]string{"fatal", "error", "warn", "info"}
	fmt.Println(unKnowns)

	unKnownCopy := [4]string{"error", "warn", "info", "fatal"}
	fmt.Println(unKnowns == unKnownCopy)
	fmt.Println("unKnownCopy:", unKnownCopy)
	fmt.Println("unKnownCopy leg", len(unKnownCopy))
	fmt.Println("unKnownCopy cap", cap(unKnownCopy))

	for unKnown := range unKnowns {
		fmt.Println("unKnowns", unKnown)
	}

	var mySlice = make([]string, 3, 6)
	fmt.Println("mySlice leg", len(mySlice))
	fmt.Println("mySlice cap", cap(mySlice))

	var slice1 = make([]string, 3, 10)
	slice1[0] = "Ram"
	slice1[1] = "Lakshman"
	slice1[2] = "Bharat"
	slice2 := append(slice1, "Satrudhan", "Luv", "Kush")
	slice3 := append(slice1, "Hanuman") // NOTE: Replaces Satrudhan in all the slices
	slice4 := append(slice1, slice2...)
	slice1[0] = "RamSita"
	fmt.Println("Slice1=============")
	for i, name := range slice1 {
		fmt.Printf("slice1[%d]=%s\n", i, name)
	}
	fmt.Println("Slice2=============")
	for i, name := range slice2 {
		fmt.Printf("slice2[%d]=%s\n", i, name)
	}
	fmt.Println("Slice3=============")
	for i, name := range slice3 {
		fmt.Printf("slice3[%d]=%s\n", i, name)
	}

	fmt.Println("Slice4=============")
	for i, name := range slice4 {
		fmt.Printf("slice4[%d]=%s\n", i, name)
	}
}
