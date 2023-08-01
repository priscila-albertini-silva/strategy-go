package internal

type Strategy interface {
	CalculateDiscount(amount float64) float64
}
