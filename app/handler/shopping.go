package handler

import (
	"awesomeProject/app/model"
	"awesomeProject/app/service"
	"fmt"
	"net/http"
	"strconv"
)

type ShoppingHandler struct {
	cs service.CartService
}

func (sh *ShoppingHandler) AddToCart(w http.ResponseWriter, req *http.Request) {
	item := req.Form.Get("item")
	quantity, _ := strconv.Atoi(req.Form.Get("quantity"))

	err := sh.cs.AddToCart(model.CartItems{
		Item:     model.Item{Name: item},
		Quantity: quantity,
	})
	if err != nil {
		fmt.Fprintf(w, "Bad Request")
	}
	fmt.Fprintf(w, "Item(s) added to cart successfully")
}

func NewShopping(cartService service.CartService) *ShoppingHandler {
	return &ShoppingHandler{cs: cartService}
}
