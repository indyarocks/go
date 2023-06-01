package main

import (
	"compositions/store"
	"github.com/fatih/color"
	"strconv"
)

func main() {
	kayak := store.NewProduct("Kayak", "WaterSports", 275)
	lifejacket := &store.Product{Name: "Life-jacket", Category: "WaterSports"}

	for _, sport := range []*store.Product{kayak, lifejacket} {
		color.Green("Name: " + sport.Name)
		color.Blue("Category: " + sport.Category)
		color.Red("Price: $" + strconv.FormatFloat(sport.Price(0.2), 'f', 2, 64))
	}
}
