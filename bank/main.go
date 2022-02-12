package main

import (
	"awesomeProject/bank/model"
	"fmt"
)

func main() {
	address1 := model.NewAddress("Chennai")
	address2 := model.NewAddress("Coimbatore")

	customer := model.Customer{}
	customer = customer.AddAccount(model.NewAccount(address1))
	customer = customer.AddAccount(model.NewAccount(address1))
	customer = customer.AddAddress(address1)

	fmt.Println("Before Update", customer)
	customer = customer.UpdateAddress(address2)
	fmt.Println("After Update", customer)
}
