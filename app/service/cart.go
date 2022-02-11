package service

import (
	"awesomeProject/app/model"
	"awesomeProject/app/repository"
	"errors"
)

type CartService interface {
	AddToCart(cartItems model.CartItems) error
}


type cartService struct {
	cartRepo repository.CartRepository
}

func (cs cartService) AddToCart(cartItems model.CartItems) error {
	cart := cs.cartRepo.Add(cartItems)
	if len(cart.CartItems) != 0 {
		return errors.New("unable to add items to cart")
	}
	return nil
}

func NewCartService(cartRepo repository.CartRepository) CartService {
	return &cartService{cartRepo: cartRepo}
}

