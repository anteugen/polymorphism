package main

import (
	"fmt"
	"github.com/anteugen/polymorphism/paymentmethods"
)

type PaymentMethod interface {
	ProcessPayment(amount float64) error
	ValidatePaymentDetails() error
}

func main() {
	payments := []PaymentMethod{
		&paymentMethods.CreditCard{"4532759734545858", "01/28", "123"},
		&paymentMethods.PayPal{"bob@example.com"},
		&paymentMethods.BankTransfer{"123456781234", "123456789", "Bob Example"},
		&paymentMethods.CreditCard{"4531759734545858", "01/28", "123"},
		&paymentMethods.PayPal{"invalid-mail.com"},  
		&paymentMethods.BankTransfer{"wrong", "123456789", "Bob Example"},        
	}

	for _, payment := range payments {
		if err := payment.ValidatePaymentDetails(); err != nil {
			fmt.Println("Error in payment details:", err)
			continue
		}
		if err := payment.ProcessPayment(100); err != nil {
			fmt.Println("Error processing payment:", err)
			continue
		}
	}

	
}
