package internal

type Customer struct {
	strategy Strategy
}

func (c *Customer) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Customer) CalculateDiscount(amount float64) float64 {
	return c.strategy.CalculateDiscount(amount)
}
