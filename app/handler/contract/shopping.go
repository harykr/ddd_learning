package contract

type AddToCartRequest struct {
	Item string `form:"item"`
	Quantity int `form:"quantity"`
}
