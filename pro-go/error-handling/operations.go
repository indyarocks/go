package main

type CategoryError struct {
	requestedCategory string
}

func (category *CategoryError) Error() string {
	return "ERROR: Category: " + category.requestedCategory + " doesn't exist."
}

type ProductCategoryTotal map[string]float64

var ProductCategoryTotalMap = make(ProductCategoryTotal)

func (slice ProductSlice) TotalCategoryPrice(category string) (total float64, err *CategoryError) {
	var categoryHasProduct bool = false
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			categoryHasProduct = true
		}
	}
	if !categoryHasProduct {
		err = &CategoryError{requestedCategory: category}
	}
	return
}

func init() {
	for _, p := range slice {
		ProductCategoryTotalMap[p.Category] += p.Price
	}
}
