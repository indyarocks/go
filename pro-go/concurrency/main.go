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

func enumerateProducts(productChannel chan<- *Product) {
	for _, product := range ProductList[:3] {
		productChannel <- product
		time.Sleep(time.Millisecond * 800)
	}
	close(productChannel)
}

func nonBlockingSendEnumerateProducts(productsChannel chan<- *Product) {
	for _, p := range ProductList {
		select {
		case productsChannel <- p:
			fmt.Println("Sent to channel:", p.Name)
		default:
			fmt.Println("Discarding Product", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(productsChannel)
}

func sendingToMultipleChannelEnumerateProducts(channel1, channel2 chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel1 <- p:
			fmt.Println("Sent to channel1", p.Name)
		case channel2 <- p:
			fmt.Println("Sent to channel2", p.Name)
		}
	}
	close(channel1)
	close(channel2)
}

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	//time.Sleep(time.Second * 5)

	dispatchChannel := make(chan DispatchNotification, 100)
	productChannel := make(chan *Product, 100)

	//var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	//var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel

	//go DispatchOrder(sendOnlyChannel)
	go DispatchOrder(chan<- DispatchNotification(dispatchChannel))
	go enumerateProducts(chan<- *Product(productChannel))
	channelCount := 2
	for {
		select {
		case order, ok := <-dispatchChannel:
			if ok {
				fmt.Println("Dispatch to:", order.Customer, ":", order.Quantity, "x", order.Product.Name)
			} else {
				fmt.Println("dispatchChannel has been closed")
				channelCount--
				dispatchChannel = nil
			}
		case product, ok := <-productChannel:
			if ok {
				fmt.Println("Product:", product.Name, "Category:", product.Category, "Price:", product.Price)
			} else {
				fmt.Println("productChannel has been closed")
				channelCount--
				productChannel = nil
			}
		default:
			if channelCount == 0 {
				goto alldone
			}
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}
	//for {
	//	select {
	//	case order, ok := <-dispatchChannel:
	//		if ok {
	//			fmt.Println("Dispatch to:", order.Customer, ":", order.Quantity, "x", order.Product.Name)
	//		} else {
	//			fmt.Println("Channel has been closed")
	//			goto alldone
	//		}
	//	default:
	//		fmt.Println("-- No message ready to be received")
	//		time.Sleep(time.Millisecond * 500)
	//	}
	//}

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
alldone:
	fmt.Println("All values received")
	testProductChannel := make(chan *Product, 5)
	go nonBlockingSendEnumerateProducts(chan<- *Product(testProductChannel))
	time.Sleep(time.Second)
	for p := range testProductChannel {
		fmt.Println("Received", p.Name)
	}

	c1 := make(chan *Product, 2)
	c2 := make(chan *Product, 1)
	go sendingToMultipleChannelEnumerateProducts(c1, c2)

	time.Sleep(time.Second)
	for p := range c1 {
		fmt.Println("Received on Channel 1", p.Name)
	}
	for p := range c2 {
		fmt.Println("Received on Channel 2", p.Name)
	}
	fmt.Println("main function complete")
}
