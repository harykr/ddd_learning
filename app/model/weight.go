package model

type Weight struct {
	Value float64
	Unit string
}

func NewWeight(value float64) Weight {
	return Weight{
		Value: value,
		Unit:  "grams",
	}
}
