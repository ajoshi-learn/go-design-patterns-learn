package factory_method

import (
	"errors"
	"fmt"
)

type PaymentType int

const (
	Cash      PaymentType = iota
	DebitCard
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(paymentType PaymentType) (PaymentMethod, error) {
	switch paymentType {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d is not recognized", paymentType))
	}
}

type CashPM struct{}
type DebitPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}
