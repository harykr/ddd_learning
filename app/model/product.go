package model

type Product struct {
	Name            string
	Price           Price
}

func NewProduct(name string, price Price) Product {
	return Product{Name: name, Price: price}
}