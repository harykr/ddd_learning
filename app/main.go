package main

import (
	"awesomeProject/app/model"
	"awesomeProject/app/model/service"
)

func main() {

	cart1 := model.NewCart()
	cart2 := model.NewCart()

	iPadPro := model.NewProduct("iPad Pro", model.InitializePrice(23000), model.NewWeight(10))
	heroPen := model.NewProduct("Hero ink Pen", model.InitializePrice(23000), model.NewWeight(20))
	cricketBat := model.NewProduct("GM Cricket bat", model.InitializePrice(23000), model.NewWeight(30))
	iPadProNew := model.NewProduct(iPadPro.Name, service.CalculateDiscountedPrice(iPadPro), model.NewWeight(41))

	cricketBatItem := model.NewItem(cricketBat, 2)

	cart1 = cart1.Add(model.NewItem(iPadProNew, 1))
	cart1 = cart1.Add(model.NewItem(heroPen, 1))
	cart1 = cart1.Add(model.NewItem(cricketBat, 2))

	cart2 = cart2.Add(cricketBatItem)
	cart2 = cart2.Add(model.NewItem(iPadProNew, 1))
	cart2 = cart2.Add(model.NewItem(heroPen, 1))
	cart2 = cart2.Remove(cricketBatItem)


	order1 := model.NewOrder()
	order1 = order1.Place(cart1)
	order1.Statement()

	order2 := model.NewOrder()
	order2 = order2.Place(cart1)
	order2.Statement()
}
