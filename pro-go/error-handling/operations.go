package main

func (slice ProductSlice) TotalCategoryPrice(category string) (total float64) {
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
		}
	}
	return
}
