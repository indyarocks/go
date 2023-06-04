package main

import "fmt"

func main() {
	var categories = []string{"Mobile", "Desktop", "Accessories", "Books"}

	for _, category := range categories {
		if total, err := slice.TotalCategoryPrice(category); err == nil {
			fmt.Println("Category:", category, "Total:", ToCurrency(total))
		} else {
			fmt.Println(err)
		}

	}
}
