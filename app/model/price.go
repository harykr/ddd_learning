package model

type Price struct {
	Value    float64
	Currency string
}

func InitializePrice(value float64) Price {
	return Price{
		Value:    value,
		Currency: "INR",
	}
}