package packageone

import "fmt"

var privateVar = "I am private"
var PublicVar = "I am public (or exported)"

func notExported() {
	fmt.Println("Inside private function")
}

func Exported() {
	fmt.Println("Inside exported function")
}
