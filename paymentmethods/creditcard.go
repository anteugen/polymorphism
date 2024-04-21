package paymentMethods

import (
	"fmt"
	"strconv"
	"unicode"
	"errors"
)

type CreditCard struct {
	CardNumber     string
	ExpirationDate string
	CVV            string
}

func reverse(str string) string {
	var result string
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func formatCardNumber(cardNumber string) string {
	var clearedCardNumber string
	for _, char := range cardNumber {
		if unicode.IsDigit(char) {
			clearedCardNumber += string(char)
		}
	}
	return clearedCardNumber
}

func luhnAlgorithm(cardNumber string) bool {
	clearedCardNumber := formatCardNumber(cardNumber)
	reversedNumber := reverse(clearedCardNumber)

	var sum int
	for i, char := range reversedNumber {
		digit, _ := strconv.Atoi(string(char))
		if i%2 == 1 {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}

func (cc *CreditCard) ValidatePaymentDetails() error {
	if len(formatCardNumber(cc.CardNumber)) != 16 {
		return errors.New("please enter a valid card number with 16 digits")
	}
	if !luhnAlgorithm(cc.CardNumber) {
		return errors.New("invalid credit card number")
	}
	return nil
}

func (cc *CreditCard) ProcessPayment(amount float64) error {
	fmt.Printf("Simulated processing of payment $%.2f\n", amount)

	return nil
}