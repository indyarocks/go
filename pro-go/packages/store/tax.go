package store

import "fmt"

const defaultTaxRate float64 = 0.2
const minThreshold = 10

type taxRate struct {
	rate, threshold float64
}

//var categoryMaxPrices = map[string]float64{
//	"Mobile":    250 + (250 * defaultTaxRate),
//	"Insurance": 150 + (150 * defaultTaxRate),
//	"Luxury":    50 + (50 * defaultTaxRate),
//}

var categoryMaxPrices = map[string]float64{
	"Mobile":    250,
	"Insurance": 150,
	"Luxury":    50,
}

func init() {
	fmt.Println("tax.go init function invoked")
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + (price * defaultTaxRate)
	}
}

// newTaxRate - Tax Rate Constructor
func newTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minThreshold {
		threshold = minThreshold
	}
	return &taxRate{rate, threshold}
}

func (taxRate *taxRate) calcPriceWithTax(p *Product) (price float64) {
	if p.price > taxRate.threshold {
		price = p.price + p.price*taxRate.rate
	} else {
		price = p.price
	}
	if max, ok := categoryMaxPrices[p.Category]; ok && price > max {
		price = max
	}
	return
}
