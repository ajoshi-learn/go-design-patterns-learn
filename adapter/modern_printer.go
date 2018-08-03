package adapter

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter == nil {
		newMsg = p.Msg
	} else {
		newMsg = p.OldPrinter.Print(p.Msg)
	}
	return
}
