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
}
