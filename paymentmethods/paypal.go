package paymentMethods

import (
	"fmt"
	"net/mail"
)

type PayPal struct {
	Email string
}

func (pp *PayPal) ValidatePaymentDetails() error {
	_, err := mail.ParseAddress(pp.Email)
    if err != nil {
		return err
	}
	return nil
}

func (pp *PayPal) ProcessPayment(amount float64) error {
	fmt.Printf("Simulated PayPal payment of $%.2f to %s\n", amount, pp.Email)
	return nil
}