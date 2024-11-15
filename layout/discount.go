package layout

import (
	"strings"

	"github.com/DoTuanAnh2k1/printing-sampa-pos/model"
)

func replaceDiscount(layout string, discountList []model.Discount) string {
	var discountOut string
	discountOut += "[DISCOUNTS]\n"
	discountLayoutFull, discountLayout := getSection(layout, `(?s)\[DISCOUNTS\](.*?)\[`)

	if len(discountList) != 0 {
		for _, discount := range discountList {
			tmp := discountLayout
			tmp = strings.ReplaceAll(tmp, "{CALCULATION NAME}", discount.Name)
			tmp = strings.ReplaceAll(tmp, "{CALCULATION DESCRIPTION}", discount.Description)
			tmp = strings.ReplaceAll(tmp, "{CALCULATION TOTAL}", discount.Total)
			tmp += "\n"
			discountOut = discountOut + tmp
		}
	}

	layout = strings.ReplaceAll(layout, discountLayoutFull, discountOut)

	return layout
}
