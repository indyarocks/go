package main

type CategoryError struct {
	requestedCategory string
}

type ProductCategoryTotal map[string]float64

var ProductCategoryTotalMap = make(ProductCategoryTotal)

type ChannelMessage struct {
	Category string
	Total    float64
	Error    *CategoryError
}

func (category *CategoryError) Error() string {
	return "ERROR: Category: " + category.requestedCategory + " doesn't exist."
}

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

func (slice ProductSlice) TotalCategoryPriceAsync(categories []string, channel chan<- ChannelMessage) {
	for _, category := range categories {
		total, err := slice.TotalCategoryPrice(category)
		channel <- ChannelMessage{
			Category: category,
			Total:    total,
			Error:    err,
		}
	}
	close(channel)
}

func init() {
	for _, p := range slice {
		ProductCategoryTotalMap[p.Category] += p.Price
	}
}
