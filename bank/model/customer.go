package model

import "github.com/google/uuid"

type Customer struct {
	Id      uuid.UUID
	Account []Account
	Address Address
}

func NewCustomer(accounts []Account, address Address) Customer {
	return Customer{
		Id:      uuid.New(),
		Account: accounts,
		Address: address,
	}
}

func (c Customer) AddAccount(account Account) Customer {
	tempAccounts := c.Account
	tempAccounts = append(tempAccounts, account)
	c.Account = tempAccounts
	return c
}

func (c Customer) AddAddress(address Address) Customer {
	c.Address = address
	return c
}

func (c Customer) UpdateAddress(address Address) Customer {
	var tempAccounts []Account
	for _, account := range c.Account {
		tempAccounts = append(tempAccounts, account.updateAddress(address))
	}
	c.Account = tempAccounts
	c.Address = address
	return c
}
