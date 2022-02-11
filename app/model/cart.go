package model

import (
	"github.com/google/uuid"
)
type Cart struct {
	Id    uuid.UUID
	Items []Item
}

func (c Cart) Add(item Item) {
	var tempItems []Item
	tempItems = append(tempItems, item)
	c.Items = tempItems
}

func (c Cart) Remove(item Item) {
	tempItems := c.Items
	var items []Item
	for _, tempItem := range tempItems {
		if tempItem != item {
			items = append(items, tempItem)
		}
	}
	c.Items = items
}

func (c Cart) Equals(cart Cart) bool {
	return c.Id == cart.Id
}