package model

import (
	"github.com/google/uuid"
)

type Cart struct {
	Id           uuid.UUID
	Items        []Item
	RemovedItems []Item
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
