package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)

	_ = keyboard.Open()
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press any key, or q to exit.")
	go listenForKeyPress()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			fmt.Println("Since you pressed Q, bye bye..")
			break
		}
		keyPressChan <- char
	}

	go doSomethingInRoutine("I'm inside do something routine")

	fmt.Println("This is first another message")

	go doSomethingInRoutine("I'm inside another something routine")

	fmt.Println("This is second another message")

	// for {
	// 	// Do nothing
	// }
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You preassed", string(key))
	}
}

func doSomethingInRoutine(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}
