package main

import (
	"fmt"
	"time"
)

func receiveDispatches(dispatchChannel <-chan DispatchNotification) {
	fmt.Println("Enumerate through channel values in for loop")
	for order := range dispatchChannel {
		fmt.Println("Dispatch to:", order.Customer, ":", order.Quantity, "x", order.Product.Name)
	}
}

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	//time.Sleep(time.Second * 5)

	dispatchChannel := make(chan DispatchNotification, 100)

	var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	//var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel

	go DispatchOrder(sendOnlyChannel)
	for {
		select {
		case order, ok := <-dispatchChannel:
			if ok {
				fmt.Println("Dispatch to:", order.Customer, ":", order.Quantity, "x", order.Product.Name)
			} else {
				fmt.Println("Channel has been closed")
				goto alldone
			}
		default:
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}
alldone:
	fmt.Println("All values received")
	//receiveDispatches(receiveOnlyChannel)
	//for {
	//	if order, open := <-dispatchChannel; open {
	//		fmt.Println("Order dispatched:", order.Customer, order.Name, order.Quantity)
	//	} else {
	//		fmt.Println("Channel has been closed", order)
	//		break
	//	}
	//}
	//fmt.Println("Enumerate through channel values in for loop")
	//for order := range dispatchChannel {
	//	fmt.Println("Dispatch to:", order.Customer, ":", order.Quantity, "x", order.Product.Name)
	//}
	//fmt.Println("Channel has been closed")

	fmt.Println("main function complete")
}
