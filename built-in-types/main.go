package main

import "fmt"

func main() {
	age := 10
	name := "Jack"
	rightHanded := true

	fmt.Printf("%s is %d years old. Right handed: %t\n", name, age, rightHanded)

	ageInTenYears := age + 10

	fmt.Printf("In 10 years, %s will be %d years old\n", name, ageInTenYears)

	isATeenager := age >= 13

	fmt.Println(name, "is a teenager:", isATeenager)

}
