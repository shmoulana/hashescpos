package layout

import (
	"strings"

	"github.com/DoTuanAnh2k1/printing-sampa-pos/model"
)

func replaceTax(layout string, taxList []model.Tax) string {
	var taxOut string
	taxOut += "[TAXES]\n"
	taxSectionLayoutFull, taxSectionLayout := getSection(layout, `(?s)\[TAXES\](.*?)\[`)
	if len(taxList) != 0 {
		for _, tax := range taxList {
			tmp := taxSectionLayout
			tmp = strings.ReplaceAll(tmp, "{TAX NAME}", tax.Name)
			tmp = strings.ReplaceAll(tmp, "{TAX RATE}", tax.Rate)
			tmp = strings.ReplaceAll(tmp, "{TAX AMOUNT}", tax.Amount)
			tmp += "\n"
			taxOut = taxOut + tmp
		}
	}

	layout = strings.ReplaceAll(layout, taxSectionLayoutFull, taxOut)
	return layout
}
