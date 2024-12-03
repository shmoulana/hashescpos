package printing

import (
	"net"
	"strings"

	"github.com/chronnie/go-escpos"
)

func PrintingWithUSBPath(data string, commandList []byte, path string) error {
	lines := strings.Split(data, "\n")
	p, err := escpos.NewUSBPrinterByPath(path)
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

func PrintingOverInternet(data string, commandList []byte, printerAddress string) error {
	conn, err := net.Dial("tcp", printerAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(commandList)
	if err != nil {
		return err
	}

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		switch line {
		case "<EB>":
			_, err := conn.Write([]byte{0x1B, 0x45, 0x01}) // ESC E 1
			if err != nil {
				return err
			}
		case "<DB>":
			_, err := conn.Write([]byte{0x1B, 0x45, 0x00}) // ESC E 0
			if err != nil {
				return err
			}
		default:
			_, err := conn.Write([]byte(line + "\n"))
			if err != nil {
				return err
			}
		}
	}

	_, err = conn.Write([]byte{0x1B, 0x64, 0x02}) // ESC d 2 (Feed 2 lines)
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte{0x1D, 0x56, 0x01}) // GS V 1 (Cut paper)
	if err != nil {
		return err
	}

	return nil
}
