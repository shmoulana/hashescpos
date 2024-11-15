package printing

import (
	"regexp"
	"strings"
)

func HandlePayment(layout string) (string, error) {
	var payment strings.Builder

	paymentSectionLayout := getPaymentSection(layout)
	paymentSection, err := handlePaymentLine(paymentSectionLayout)
	if err != nil {
		return "", err
	}
	payment.WriteString(paymentSection)

	return payment.String(), nil
}

func getPaymentSection(layout string) (paymentSectionLayout string) {
	// get payment section
	rePayment := regexp.MustCompile(`(?s)\[PAYMENTS\](.*?)\[`)
	matchesPayment := rePayment.FindStringSubmatch(layout)
	if len(matchesPayment) > 1 {
		paymentSectionLayout = matchesPayment[1]
	}
	return
}

func handlePaymentLine(paymentLine string) (string, error) {
	var payment strings.Builder
	lines := strings.Split(paymentLine, "\n")
	for _, line := range lines {
		lineFormat, err := GetLineFormat(line)
		if err != nil {
			return payment.String(), err
		}
		payment.WriteString(lineFormat)
	}
	return removeEmptyLines(payment.String()), nil
}
