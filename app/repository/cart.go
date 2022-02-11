package repository

import "awesomeProject/app/model"

type CartRepository interface {
	Add(cartItem model.CartItems) model.Cart
}

type cartRepository struct {
	cart model.Cart
}

func (sr *cartRepository) Add(cartItem model.CartItems) model.Cart {
	sr.cart.Add(cartItem)
	return sr.cart
}

func NewCartRepository() CartRepository {
	return &cartRepository{}
}
