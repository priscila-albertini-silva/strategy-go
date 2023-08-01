package internal

type RegularCustomerStrategy struct{}

func (rcs *RegularCustomerStrategy) CalculateDiscount(amount float64) float64 {
	return amount * 0.1
}
