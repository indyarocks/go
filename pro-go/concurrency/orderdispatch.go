package main

import (
	"fmt"
	"math/rand"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"Chandan", "Darsh", "Tithi", "Sheenu"}

func DispatchOrder(channel chan<- DispatchNotification) {
	var orderCount = rand.Intn(10) + 2
	fmt.Println("Order Count:", orderCount)
	for i := 0; i < orderCount; i++ {
		customer := Customers[rand.Intn(len(Customers)-1)]
		quantity := rand.Intn(10) + 1
		channel <- DispatchNotification{
			Customer: customer,
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
			Quantity: quantity,
		}
		//if i == 1 {
		//	notification := <-channel
		//	fmt.Println("Read:", notification.Customer)
		//}
	}
	close(channel)
}
