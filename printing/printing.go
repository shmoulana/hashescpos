package printing

import (
	"strings"

	// _ "github.com/chronnie/go-escposv"
	"github.com/chronnie/go-escpos"
)

func Printing(data string, commandList []byte) error {
	lines := strings.Split(data, "\n")
	p, err := escpos.NewUSBPrinterByPath("")
	if err != nil {
		return err
	}

	err = p.Init()
	if err != nil {
		return err
	}

	for _, command := range commandList {
		p.Write(string(command))
	}

	for _, line := range lines {
		if line == "<EB>" {
			p.EnableBold()
		} else if line == "<DB>" {
			p.DisableBold()
		}
		p.Print(line)
	}

	p.Feed(2)
	p.Cut()
	p.End()
	return nil
}
