package model

import (
	"github.com/bradfitz/iter"
	"github.com/google/uuid"
)

type Order struct {
	Id       uuid.UUID
	Products []Product
}

func (o Order) Place(cart Cart) Order {
	var products []Product
	for _, item := range cart.Items {
		for _ = range iter.N(item.Quantity) {
			products = append(products, item.Product)
		}
	}
	o.Products = products
	return o
}