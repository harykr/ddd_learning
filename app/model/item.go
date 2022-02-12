package model

import "github.com/google/uuid"

type Item struct {
	Id       uuid.UUID
	Product  Product
	Quantity int
}

func NewItem(product Product, quantity int) Item {
	return Item{
		Id:       uuid.New(),
		Product:  product,
		Quantity: quantity,
	}
}
