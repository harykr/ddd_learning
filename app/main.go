package main

import (
	"awesomeProject/app/model"
	"awesomeProject/app/model/service"
	"fmt"
)

func main() {

	cart1 := model.NewCart()
	cart2 := model.NewCart()

	iPadPro := model.NewProduct("iPad Pro", model.InitializePrice(23000, "INR"))
	heroPen := model.NewProduct("Hero ink Pen", model.InitializePrice(23000, "INR"))
	cricketBat := model.NewProduct("GM Cricket bat", model.InitializePrice(23000, "INR"))
	iPadProNew := model.NewProduct(iPadPro.Name, service.CalculateDiscountedPrice(iPadPro))

	cricketBatItem := model.NewItem(cricketBat, 2)

	cart1 = cart1.Add(model.NewItem(iPadProNew, 1))
	cart1 = cart1.Add(model.NewItem(heroPen, 1))
	cart1 = cart1.Add(model.NewItem(cricketBat, 2))

	fmt.Println("Cart1 Items", cart1.Items)
	fmt.Println("Cart1 Items", cart1.RemovedItems)

	cart2 = cart2.Add(cricketBatItem)
	cart2 = cart2.Add(model.NewItem(iPadProNew, 1))
	cart2 = cart2.Add(model.NewItem(heroPen, 1))
	cart2 = cart2.Remove(cricketBatItem)

	fmt.Println("Cart2 Items", cart2.Items)
	fmt.Println("Cart2 Items", cart2.RemovedItems)

	order1 := model.NewOrder()
	order1 = order1.Place(cart1)

	fmt.Println("Order for Cart 1", order1)

}
