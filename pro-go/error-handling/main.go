package main

import "fmt"

func main() {
	var categories = []string{"Mobile", "Desktop", "Accessories", "Books", "SelfCare"}

	for _, category := range categories {
		if total, err := slice.TotalCategoryPrice(category); err == nil {
			fmt.Println("Category:", category, "Total:", ToCurrency(total))
		} else {
			fmt.Println(err)
		}

	}
	// Using map
	for _, category := range categories {
		if total, ok := ProductCategoryTotalMap[category]; ok {
			fmt.Println("Category:", category, "Total:", ToCurrency(total))
		} else {
			fmt.Println(&CategoryError{requestedCategory: category})
		}
	}

}
