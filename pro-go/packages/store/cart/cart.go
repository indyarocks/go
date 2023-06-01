package cart

import (
	"fmt"
	"packages/store"
)

type Cart struct {
	CustomerName string
	Product      []store.Product
}

func init() {
	fmt.Println("cart.go init function invoked")
}

func (cart *Cart) GetTotal() (total float64) {
	for _, product := range cart.Product {
		total += product.Price()
	}
	return
}
