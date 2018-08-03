package adapter

import (
	"testing"
	"fmt"
)

func TestAdapter(t *testing.T) {
	msg := "Hello world"
	adapter := PrinterAdapter{&MyLegacyPrinter{}, msg}
	returnedMessage := adapter.PrintStored()
	fmt.Println(returnedMessage)
	if returnedMessage != "Legacy Printer: Hello world\n" {
		t.Errorf("Message didn't match: %s\n", returnedMessage)
	}
}

func TestNilAdapter(t *testing.T) {
	msg := "Hello world"
	adapter := PrinterAdapter{nil, msg}
	returnedMessage := adapter.PrintStored()
	if returnedMessage != "Hello world\n" {
		t.Errorf("Message didn't match: %s\n", returnedMessage)
	}
}