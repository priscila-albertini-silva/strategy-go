package internal

type VIPCustomerStrategy struct{}

func (vcs *VIPCustomerStrategy) CalculateDiscount(amount float64) float64 {
	return amount * 0.2
}
