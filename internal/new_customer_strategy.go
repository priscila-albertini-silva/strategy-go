package internal

type NewCustomerStrategy struct{}

func (ncs *NewCustomerStrategy) CalculateDiscount(amount float64) float64 {
	return 0
}
