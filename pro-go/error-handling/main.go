package main

import "fmt"

func main() {
	var categories = []string{"Mobile", "Desktop", "Accessories", "Test"}

	for _, category := range categories {
		fmt.Println("Category:", category, "Total:", slice.TotalCategoryPrice(category))
	}
}
