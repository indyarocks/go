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

	fmt.Println("Receive total and error via channel=====")
	totalChannel := make(chan ChannelMessage)
	go slice.TotalCategoryPriceAsync(categories, chan<- ChannelMessage(totalChannel))

	for message := range totalChannel {
		if message.Error == nil {
			fmt.Println("Category:", message.Category, "Total:", message.Total)
		} else {
			fmt.Println("ERROR", message.Error)
		}
	}

	fmt.Println("\n\nReceive total and generic error via channel=====")
	totalChannelGenericError := make(chan ChannelMessageWithGenericError)
	go slice.TotalCategoryPriceGenericErrorAsync(categories, totalChannelGenericError)
	for message := range totalChannelGenericError {
		if message.Error == nil {
			fmt.Println("Category:", message.Category, "Total:", message.Total)
		} else {
			fmt.Println("ERROR", message.Error)
		}
	}
}
