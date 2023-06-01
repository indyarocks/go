package main

import (
	"compositions/store"
	"fmt"
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

	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}
	for _, boat := range boats {
		fmt.Println("Conventional:", boat.Product.Name, "Direct:", boat.Name)
		fmt.Println("Boat:", boat.Name, "Price:", boat.Price(0.2))
	}

	rentalBoats := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "Chandan", "Darsh"),
		store.NewRentalBoat("Yatch", 50000, 5, true, true, "Shweta", "Tithi"),
		store.NewRentalBoat("Super Yatch", 100000, 15, true, true, "Mom", "Dad"),
	}

	for _, r := range rentalBoats {
		fmt.Println("Rental Boat", r.Name, "Crewed", r.IncludeCrew, "Rental Price", r.Price(0.2))
		color.Red("Captain:" + r.Captain)
		color.Green("First Officer" + r.FirstOfficer)
	}

	deal := store.NewSpecialDeal("Weekend Special", kayak, 50)
	dealName, dealPrice, _ := deal.GetDetails()
	color.Green("Deal Name:" + dealName)
	color.Green("Product Name:" + deal.Product.Name)
	color.Red("Original Product Price:" + strconv.FormatFloat(deal.Product.Price(0), 'f', 2, 64))
	color.Green("Discounted Price:" + strconv.FormatFloat(dealPrice, 'f', 2, 64))

	//	Promotion Ambiguity - Price method is defined on both SpecialDeal and Product
	type OfferBundle struct {
		*store.SpecialDeal
		*store.Product
	}

	bundle := OfferBundle{
		store.NewSpecialDeal("Weekend Special", kayak, 50),
		store.NewProduct("LifeJacket", "WaterSports", 100),
	}

	//fmt.Println("Price:", bundle.Price(0))
	fmt.Println("Special Deal Price:", bundle.SpecialDeal.Price(0))
	fmt.Println("Product Price:", bundle.Product.Price(0))

	products := map[string]store.ItemForSale{
		"iPhone":      store.NewProduct("iPhone", "Mobile", 1299),
		"Kayak":       store.NewBoat("Kayak", 279, 1, false),
		"specialDeal": store.NewSpecialDeal("Weekend Sales", kayak, 50),
		"rentedYatch": store.NewRentalBoat("Yatch", 100000, 15, true, true, "Chandan", "Darsh"),
	}
	for key, p := range products {
		fmt.Println("Key", key, "Price:", p.Price(0.2))
	}
}
