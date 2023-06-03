package main

import (
	"fmt"
	"math/rand"
)

type DispatchNotification struct {
	CustomerName string
	*Product
	Quantity int
}

var Customers = []string{"Chandan", "Darsh", "Tithi", "Sheenu"}

func DispatchOrder(channel chan DispatchNotification) {
	var orderCount = rand.Intn(20) + 2
	fmt.Println("Order Count:", orderCount)
	for i := 0; i < orderCount; i++ {
		customer := Customers[rand.Intn(len(Customers)-1)]
		quantity := rand.Intn(10) + 1
		channel <- DispatchNotification{
			CustomerName: customer,
			Product:      ProductList[rand.Intn(len(ProductList)-1)],
			Quantity:     quantity,
		}
	}
	close(channel)
}
