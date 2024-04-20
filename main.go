package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type PaymentMethod interface {
	ProcessPayment(amount float64) error
	ValidatePaymentDetails() error
}

type CreditCard struct {
	CardNumber     string
	ExpirationDate string
	CVV            string
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func LuhnAlgorithm(cardNumber string) (bool, string) {
	var clearedCardNumber string

	for _, char := range cardNumber {
		if unicode.IsDigit(char) {
			clearedCardNumber += string(char)
		}
	}

	if len(clearedCardNumber) != 16 {
		return false, "Error: please enter a card number with 16 digits"
	}

	reversedNumber := reverse(clearedCardNumber)

	var sum int
	for i, char := range reversedNumber {
		digit, _ := strconv.Atoi(string(char))
		if i%2 == 0 {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	if sum%10 == 0 {
		return true, "Card number is valid."
	}
	return false, "Card number is invalid."
}

func main() {
	valid, message := LuhnAlgorithm("4532 7597 3454 5858")
	fmt.Println(valid, message)
	valid2, message2 := LuhnAlgorithm("5555 5555 5555 4444 ")
	fmt.Println(valid2, message2)
}
