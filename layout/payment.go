package layout

import (
	"strings"

	"github.com/DoTuanAnh2k1/printing-sampa-pos/model"
)

func replacePayment(layout string, paymentList []model.Payment) string {
	var paymentOut string
	paymentOut += "[PAYMENTS]\n"
	paymentSectionLayoutFull, paymentSectionLayout := getSection(layout, `(?s)\[PAYMENTS\](.*?)\[`)

	if len(paymentList) != 0 {
		for _, payment := range paymentList {
			tmp := paymentSectionLayout
			tmp = strings.ReplaceAll(tmp, "{PAYMENT NAME}", payment.Name)
			tmp = strings.ReplaceAll(tmp, "{PAYMENT AMOUNT}", payment.Amount)
			tmp += "\n"
			paymentOut = paymentOut + tmp
		}
	}

	layout = strings.ReplaceAll(layout, paymentSectionLayoutFull, paymentOut)
	return layout
}
