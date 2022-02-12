package model

import "github.com/google/uuid"

type Item struct {
	Id       uuid.UUID
	Product  Product
	Quantity int
}
