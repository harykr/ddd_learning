package model

type Product struct {
	Name   string
	Price  Price
	Weight Weight
}

func NewProduct(name string, price Price, weight Weight) Product {
	return Product{Name: name, Price: price, Weight: weight}
}
