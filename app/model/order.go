package model

import (
	"fmt"
	"github.com/google/uuid"
)

type TotalCost struct {
	ShippingCost Price
	ProductCost  Price
}

type Order struct {
	Id        uuid.UUID
	Products  []Product
	TotalCost TotalCost
	Amount    Price
}

func NewOrder() Order {
	return Order{Id: uuid.New()}
}

func (o Order) Statement() {
	fmt.Println("********************************")
	fmt.Println(fmt.Sprintf("Order Id : %s", o.Id.String()))
	fmt.Println(fmt.Sprintf("Shipping Cost : %.2f (%s)", o.TotalCost.ShippingCost.Value, o.TotalCost.ShippingCost.Currency))
	fmt.Println(fmt.Sprintf("Product Cost : %.2f (%s)", o.TotalCost.ProductCost.Value, o.TotalCost.ProductCost.Currency))
	fmt.Println(fmt.Sprintf("Total Cost : %.2f (%s)", o.Amount.Value, o.Amount.Currency))
	fmt.Println("********************************")
}

func (o Order) Place(cart Cart) Order {
	o.Products = cart.Products()
	cart.CheckOut()
	o.TotalCost.ShippingCost = o.CalculateShippingCost()
	o.TotalCost.ProductCost = o.CalculateProductCost()
	o.Amount = o.CalculateTotalCost()
	return o
}

func (o Order) CalculateTotalCost() Price {
	return InitializePrice(o.TotalCost.ProductCost.Value + o.TotalCost.ShippingCost.Value)
}

func (o Order) CalculateShippingCost() Price {
	var shippingWeight float64
	for _, product := range o.Products {
		shippingWeight += product.Weight.Value
	}
	return InitializePrice(shippingWeight * 0.01)
}

func (o Order) CalculateProductCost() Price {
	var totalPrice float64
	for _, product := range o.Products {
		totalPrice += product.Price.Value
	}
	return InitializePrice(totalPrice)
}
