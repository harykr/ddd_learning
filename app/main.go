package main

import (
	"awesomeProject/app/model"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	cart1 := model.Cart{Id: uuid.New()}
	cart1.Add(model.Item{
		Name:     "IPad Pro",
		Quantity: 1,
	})

	cart1.Add(model.Item{
		Name:     "Hero ink Pen",
		Quantity: 1,
	})

	cart1.Add(model.Item{
		Name:     "Hero ink Pen",
		Quantity: 1,
	})

	cart1.Add(model.Item{
		Name:     "GM Cricket bat",
		Quantity: 2,
	})

	cart1.Remove(model.Item{
		Name:     "Hero ink Pen",
		Quantity: 1,
	})

	cart2 := model.Cart{Id: uuid.New()}
	cart2.Add(model.Item{
		Name:     "IPad Pro",
		Quantity: 1,
	})

	cart2.Add(model.Item{
		Name:     "Hero ink Pen",
		Quantity: 1,
	})

	cart2.Add(model.Item{
		Name:     "Hero ink Pen",
		Quantity: 1,
	})

	cart2.Add(model.Item{
		Name:     "GM Cricket bat",
		Quantity: 2,
	})

	cart2.Remove(model.Item{
		Name:     "Hero ink Pen",
		Quantity: 1,
	})

	fmt.Println(cart1.Equals(cart2))
}
