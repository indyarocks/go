package main

import "fmt"

// Interface
type Animal interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func main() {
	dog := Dog{
		Name:         "Dog",
		Sound:        "Woof",
		NumberOfLegs: 4,
	}

	cat := Cat{
		Name:         "Cat",
		Sound:        "Meow",
		NumberOfLegs: 4,
		HasTail:      true,
	}

	Riddle(&dog)
	Riddle(&cat)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says "%s" and has "%d" legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
