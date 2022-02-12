package model

import "github.com/google/uuid"

type Account struct {
	Id      uuid.UUID
	Address Address
}

func NewAccount(address Address) Account {
	return Account{Id: uuid.New(), Address: address}
}

func (a Account) updateAddress(address Address) Account {
	a.Address = address
	return a
}
