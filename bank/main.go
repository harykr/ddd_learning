package main

import (
	"awesomeProject/bank/model"
	"fmt"
)

func main() {
	var accounts []model.Account
	address1 := model.NewAddress("Chennai")
	address2 := model.NewAddress("Coimbatorw")

	accounts = append(accounts, model.NewAccount(address1))
	accounts = append(accounts, model.NewAccount(address1))
	customer := model.NewCustomer(accounts, address1)
	fmt.Println("Before Update", customer)

	customer = customer.UpdateAddress(address2)
	fmt.Println("After Update", customer)
}
