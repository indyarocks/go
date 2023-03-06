package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var firstNumber = 2
	// var secondNumber = 5
	// var subtraction = 7
	// var answer int

	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instructions
	fmt.Println("Guess the Number game")
	fmt.Println("----------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10 and press ENTER when ready")
	reader.ReadString('\n')

	// take them through the game

	// give them the answer

}
