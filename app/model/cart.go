package model

import (
	"github.com/bradfitz/iter"
	"github.com/google/uuid"
)

type Cart struct {
	Id           uuid.UUID
	Items        []Item
	RemovedItems []Item
	IsCheckedOut bool
}

func NewCart() Cart {
	return Cart{Id: uuid.New()}
}

func (c Cart) Add(item Item) Cart {
	tempItems := c.Items
	tempItems = append(tempItems, item)
	c.Items = tempItems
	return c
}

func (c Cart) Remove(item Item) Cart {
	var items, removedItems []Item
	for _, tempItem := range c.Items {
		if tempItem.Id != item.Id {
			items = append(items, tempItem)
		}
	}
	removedItems = append(removedItems, item)

	c.RemovedItems = removedItems
	c.Items = items
	return c
}

func (c Cart) Equals(cart Cart) bool {
	return c.Id == cart.Id
}

func (c Cart) CheckOut() {
	c.IsCheckedOut = true
}

func (c Cart) Products() ([]Product) {
	var products []Product
	for _, item := range c.Items {
		for _ = range iter.N(item.Quantity) {
			products = append(products, item.Product)
		}
	}
	return products
}