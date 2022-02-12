package main

import (
	"awesomeProject/app/model"
	"awesomeProject/app/model/service"
	"fmt"
	"github.com/google/uuid"
)

func main() {

	cart1 := model.Cart{Id: uuid.New()}
	cart2 := model.Cart{Id: uuid.New()}

	iPadPro := model.Product{Name: "iPad Pro", Price: model.Price{Value: 23000, Currency: "INR"}}
	heroPen := model.Product{Name: "Hero ink Pen", Price: model.Price{Value: 1300, Currency: "INR"}}
	cricketBat := model.Product{Name: "GM Cricket bat", Price: model.Price{Value: 2300, Currency: "INR"}}


	iPadProNew := iPadPro
	iPadProNew = service.CalculateDiscountedPrice(iPadProNew, iPadPro)

	cart1 = cart1.Add(model.Item{
		Id:       uuid.New(),
		Product:  iPadProNew,
		Quantity: 1,
	})

	cart1 = cart1.Add(model.Item{
		Id:       uuid.New(),
		Product:  heroPen,
		Quantity: 1,
	})

	cart1 = cart1.Add(model.Item{
		Id:       uuid.New(),
		Product:  cricketBat,
		Quantity: 2,
	})

	fmt.Println("Cart1 Items", cart1.Items)

	cart2 = cart2.Add(model.Item{
		Id:       uuid.New(),
		Product:  iPadPro,
		Quantity: 1,
	})

	cart2heroPenId := uuid.New()
	cart2 = cart2.Add(model.Item{
		Id:       cart2heroPenId,
		Product:  heroPen,
		Quantity: 1,
	})

	cart2 = cart2.Add(model.Item{
		Id:       uuid.New(),
		Product:  cricketBat,
		Quantity: 2,
	})

	fmt.Println("Cart2 Items Before Removal", cart2.Items)
	cart2 = cart2.Remove(model.Item{
		Id:       cart2heroPenId,
		Product:  heroPen,
		Quantity: 1,
	})
	fmt.Println("Cart2 Items After Removal", cart2.Items)
	fmt.Println("Removed Items", cart2.RemovedItems)
	fmt.Println("Cart Equals", cart1.Equals(cart2))

}
