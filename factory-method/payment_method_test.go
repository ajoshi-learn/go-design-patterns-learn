package factory_method

import (
	"testing"
	"strings"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	paymentMethod, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method 'Cash' shoud exist")
	}

	msg := paymentMethod.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodDebit(t *testing.T) {
	paymentMethod, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment method 'DebitCard' shoud exist")
	}

	msg := paymentMethod.Pay(10.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNotExistent(t *testing.T) {
	_, err := GetPaymentMethod(10)
	if err == nil {
		t.Error("A payment method 10 must return an error")
	}
	t.Log("LOG:", err)
}