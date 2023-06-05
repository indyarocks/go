package main

import "fmt"

func main() {
	defer func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Recovered", err.Error())
				//panic(err) // NOTE Raising panic inside recovery
			} else if str, ok := arg.(string); ok {
				fmt.Println("Recovered", str)
			} else {
				fmt.Println("Error recovered")
			}
		}
	}()
	var categories = []string{"Mobile", "Desktop", "Accessories", "Books", "SelfCare"}

	for _, category := range categories {
		if total, err := slice.TotalCategoryPrice(category); err == nil {
			fmt.Println("Category:", category, "Total:", ToCurrency(total))
		} else {
			fmt.Println(err)
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

	fmt.Println("Using map==========")
	for _, category := range categories {
		if total, ok := ProductCategoryTotalMap[category]; ok {
			fmt.Println("Category:", category, "Total:", ToCurrency(total))
		} else {
			// NOTE: As CategoryError is a struct implementing Error interface
			// Thus &CategoryError{requestedCategory: category} will print the error message
			// i.e. returned string of Error() function
			fmt.Println(&CategoryError{requestedCategory: category})
			//panic(&CategoryError{requestedCategory: category})
		}
	}

	fmt.Println("Check the recovery inside the goroutine ======")
	recoveryChannel := make(chan CategoryCountMessage, 10)
	go ProcessCategories(categories, chan<- CategoryCountMessage(recoveryChannel))
	for msg := range recoveryChannel {
		if msg.TerminalError == nil {
			fmt.Println("Category:", msg.Category, "Total Count", msg.Count)
		} else {
			fmt.Println("####A terminal error occurred####")
		}

	}
}
