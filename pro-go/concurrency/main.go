package main

import (
	"fmt"
)

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	//time.Sleep(time.Second * 5)

	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrder(dispatchChannel)
	for {
		if order, open := <-dispatchChannel; open {
			fmt.Println("Order dispatched:", order.CustomerName, order.Name, order.Quantity)
		} else {
			fmt.Println("Channel has been closed", order)
			break
		}
	}
	fmt.Println("main function complete")
}
