package main

import (
	"fmt"

	"github.com/priscila-albertini/strategy/internal"
)

func main() {
	regularCustomer := &internal.Customer{}
	regularCustomer.SetStrategy(&internal.RegularCustomerStrategy{})

	vipCustomer := &internal.Customer{}
	vipCustomer.SetStrategy(&internal.VIPCustomerStrategy{})

	newCustomer := &internal.Customer{}
	newCustomer.SetStrategy(&internal.NewCustomerStrategy{})

	amount := 100.0

	regularCustomerAmount := regularCustomer.CalculateDiscount(amount)
	vipCustomerAmount := vipCustomer.CalculateDiscount(amount)
	newCustomerAmount := newCustomer.CalculateDiscount(amount)

	fmt.Printf("Cliente regular, valor a pagar: R$%.2f\n", regularCustomerAmount)
	fmt.Printf("Cliente VIP, valor a pagar: R$%.2f\n", vipCustomerAmount)
	fmt.Printf("Novo cliente, valor a pagar: R$%.2f\n", newCustomerAmount)
}
