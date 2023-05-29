package main

import "fmt"

type calcFunc func(float64) float64

type Product struct {
	name  string
	price float64
}

var giveAwayGift = false

//func calcWithTax(price float64) float64 {
//	return price * 0.1
//}
//
//func calcWithoutTax(price float64) float64 {
//	return 0
//}

func printPrice(product string, price float64, priceCalc calcFunc) {
	fmt.Println("====Product", product, "Final Price", priceCalc(price))
}

func selectCalculator(price float64) calcFunc {
	if price > 300 {
		//return calcWithTax
		return func(price float64) float64 {
			return price * 0.1
		}
	} else {
		//return calcWithoutTax
		return func(price float64) float64 {
			return 0
		}
	}
}

func priceCalcFactory(threshold float64, rate float64) calcFunc {
	fixedPrizeGiveAway := giveAwayGift
	return func(price float64) float64 {
		if fixedPrizeGiveAway {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		} else {
			return price
		}
	}
}

func main() {
	products := map[string]float64{
		"iPhone":  1499,
		"OnePlus": 699,
		"Mi":      299,
	}
	for product, price := range products {
		var calTax func(float64) float64
		var calTaxUsingSelector func(float64) float64 = selectCalculator(price)

		fmt.Println("function calTax assigned", calTax != nil)
		//if price > 300 {
		//	calTax = calcWithTax
		//} else {
		//	calTax = calcWithoutTax
		//}

		//fmt.Println("function calTax assigned", calTax != nil)
		//tax := calTax(price)
		//fmt.Println("product", product, "Tax amount", tax)
		//printPrice(product, price, calTax)
		printPrice(product, price, calTaxUsingSelector)
	}

	waterSportsProducts := map[string]float64{
		"ball":         100,
		"safetyJacket": 200,
		"glass":        50,
		"cap":          50,
	}
	badmintonSportsProducts := map[string]float64{
		"shuttle": 20,
		"racket":  200,
		"net":     150,
		"shoes":   400,
	}
	giveAwayGift = true
	waterSportCalc := priceCalcFactory(100, 0.2)
	giveAwayGift = false
	bandmintonCalc := priceCalcFactory(150, 0.1)

	fmt.Println("====== Water Sports ======")
	for name, price := range waterSportsProducts {
		printPrice(name, price, waterSportCalc)
	}
	fmt.Println("====== Badminton Sports ======")

	for name, price := range badmintonSportsProducts {
		printPrice(name, price, bandmintonCalc)
	}
}
