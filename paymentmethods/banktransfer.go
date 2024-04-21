package paymentMethods

import (
	"fmt"
	"errors"
	"regexp"
)

type BankTransfer struct {
	AccountNumber string
	RoutingNumber string
	AccountName   string
}

func (bt *BankTransfer) ValidatePaymentDetails() error {
	if matched, _ := regexp.MatchString(`^\d{8,12}$`, bt.AccountNumber); !matched {
		return errors.New("invalid account number")
	}
	if matched, _ := regexp.MatchString(`^\d{9}$`, bt.RoutingNumber); !matched {
		return errors.New("invalid routing number")
	}
	return nil
}

func (bt *BankTransfer) ProcessPayment(amount float64) error {
	fmt.Printf("Processing direct bank transfer of $%.2f from account %s at bank %s.\n", amount, bt.AccountNumber, bt.RoutingNumber)
	return nil
}