package pos

import (
	"errors"

	layout_utils "github.com/DoTuanAnh2k1/printing-sampa-pos/layout"
	"github.com/DoTuanAnh2k1/printing-sampa-pos/printing"
	"github.com/DoTuanAnh2k1/printing-sampa-pos/utils"
)

func (p *Pos) PrintWithUSB(usbPath string) error {
	layout, err := utils.ReadFromFile(p.layoutPath)
	if err != nil {
		return err
	}
	layoutFill := layout_utils.FillValueLayout(p.ticket, layout)

	layoutData, commandList, err := printing.HandleLayout(layoutFill)
	if err != nil {
		return err
	}

	err = printing.PrintingWithUSBPath(layoutData, commandList, usbPath)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pos) PrintWithInternet(addr, port string) error {
	if addr == "" {
		return errors.New("missing printing's address")
	}

	if port == "" {
		return errors.New("missing printing's port")
	}

	layout, err := utils.ReadFromFile(p.layoutPath)
	if err != nil {
		return err
	}
	layoutFill := layout_utils.FillValueLayout(p.ticket, layout)

	layoutData, commandList, err := printing.HandleLayout(layoutFill)
	if err != nil {
		return err
	}

	host := addr + ":" + port
	err = printing.PrintingOverInternet(layoutData, commandList, host)
	if err != nil {
		return err
	}
	return nil
}
