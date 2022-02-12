package model

type Address struct {
	City string
}

func NewAddress(city string) Address {
	return Address{City: city}
}
