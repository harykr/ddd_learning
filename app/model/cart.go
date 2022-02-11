package model

type Cart struct {
	CartItems      []CartItems
}

func (c Cart) Add(item CartItems) {
	var cartItems []CartItems
	cartItems = append(cartItems, item)
	c.CartItems = cartItems
}
