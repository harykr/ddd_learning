package model

type Account struct {
	Address Address
}

func NewAccount(address Address) Account {
	return Account{Address: address}
}

func (a Account) updateAddress(address Address) Account {
	a.Address = address
	return a
}