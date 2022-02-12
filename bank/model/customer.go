package model

type Customer struct {
	Account []Account
	Address Address
}

func NewCustomer(accounts []Account, address Address) Customer {
	return Customer{
		Account: accounts,
		Address: address,
	}
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