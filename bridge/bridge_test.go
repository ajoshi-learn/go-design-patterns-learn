package bridge

import (
	"testing"
	"errors"
)

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("content is empty")
	return
}

func TestPrinterImpl1(t *testing.T) {
	api1 := PrinterImpl1{}

	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPrinterImpl2(t *testing.T) {
	testWriter := TestWriter{}
	api2 := PrinterImpl2{&testWriter}

	expectedMessage := "Hello"
	err := api2.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("Error trying to use API2 implementation: %s\n", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly to the io.Writer.\n Expected: %s\n Actual: %s\n", expectedMessage, testWriter.Msg)
	}
}

func TestNormalPrinter(t *testing.T) {
	expectedMessage := "Hello io.Writer"

	normal := NormalPrinter{expectedMessage, &PrinterImpl1{}}
	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPackPrinter(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packt: " + passedMessage

	testWriter := TestWriter{}
	packt := PacktPrinter{passedMessage, &PrinterImpl2{&testWriter}}
	err := packt.Print()
	if err != nil {
		t.Errorf(err.Error())
	}
	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message in io.Writer doesn't macth actual.\n Actual: %s\n Expected: %s\n", testWriter.Msg, expectedMessage)
	}
}
