package service

import (
	"awesomeProject/app/model"
	"os"
	"strconv"
)

func CalculateDiscountedPrice(product model.Product, competitorProduct model.Product) model.Product {
	discountRate, _ := strconv.ParseFloat(os.Getenv("DISCOUNT_RATE"), 64)
	if discountRate == 0.0 {
		discountRate = 0.9
	} else {
		discountRate = 1 - (discountRate/100)
	}
	product.DiscountedPrice = model.Price{Value: discountRate * competitorProduct.Price.Value, Currency: "INR"}
	return product
}
