package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// Says - Receiving function
func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs\n", a.Name, a.NumberOfLegs)
}

func main() {
	cat := Animal{
		Name:         "Cat",
		Sound:        "Meow",
		NumberOfLegs: 4,
	}

	cat.Says()
	cat.HowManyLegs()

	dog := Animal{
		Name:         "Dog",
		Sound:        "Woof",
		NumberOfLegs: 4,
	}

	dog.Says()
	dog.HowManyLegs()

	z := addTwoNumbers(2, 4)
	fmt.Println(z)
	numbers := getNumbers("Please enter number you want to add. Press any other key to end.")
	fmt.Println(numbers)

	myTotal := sumMany(1, 2, 3, 4, 5)
	fmt.Println("Total of 1, 2, 3, 4, 5 is", myTotal)

	anotherTotal := sumMany(numbers...)
	fmt.Println("Total of ", numbers, "is", anotherTotal)
}

func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}
	return total
}

func addTwoNumbers(x, y int) (sum int) {
	sum = x + y
	return
}

func getNumbers(s string) []int {
	fmt.Println(s)
	fmt.Println("==>")
	var numbers []int
	reader = bufio.NewReader(os.Stdin)
	for {
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Error", err)
			break
		} else {
			numbers = append(numbers, num)
		}
	}
	return numbers
}
