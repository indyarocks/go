package data

import "fmt"

func init() {
	fmt.Println("data.go init function invoked")
}

func GetData() []string {
	return []string{"Chandan", "Tithi", "Darsh"}
}
