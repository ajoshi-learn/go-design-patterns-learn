package bridge

import (
	"io"
	"fmt"
	"errors"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("you need to pass io.Writer")
	}
	fmt.Fprintf(p.Writer, "%s", msg)
	return nil
}
