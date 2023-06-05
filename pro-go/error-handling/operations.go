package main

import (
	"errors"
	"fmt"
)

type CategoryError struct {
	requestedCategory string
}

func (category *CategoryError) Error() string {
	return "ERROR: Category: " + category.requestedCategory + " doesn't exist."
}

type ProductCategoryTotal map[string]float64

var ProductCategoryTotalMap = make(ProductCategoryTotal)

type ChannelMessage struct {
	Category string
	Total    float64
	Error    *CategoryError
}

type ChannelMessageWithGenericError struct {
	Category string
	Total    float64
	Error    error
}

func (slice ProductSlice) TotalCategoryPriceWithoutExplicitErrorStruct(category string) (total float64, err error) {
	var categoryHasProduct bool = false
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			categoryHasProduct = true
		}
	}
	if !categoryHasProduct {
		err = errors.New("can't find category")
		err = fmt.Errorf("can't find category: %v", category)
	}
	return
}

func (slice ProductSlice) TotalCategoryPriceGenericErrorAsync(categories []string, channel chan<- ChannelMessageWithGenericError) {
	for _, category := range categories {
		total, err := slice.TotalCategoryPriceWithoutExplicitErrorStruct(category)
		channel <- ChannelMessageWithGenericError{
			Category: category,
			Total:    total,
			Error:    err,
		}
	}
	close(channel)
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
