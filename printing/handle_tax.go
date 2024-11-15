package printing

import (
	"regexp"
	"strings"
)

func HandleTax(layout string) (string, error) {
	var Tax strings.Builder

	taxSectionLayout := getTaxSection(layout)
	taxSection, err := handleTaxLine(taxSectionLayout)
	if err != nil {
		return "", err
	}
	Tax.WriteString(taxSection)

	return Tax.String(), nil
}

func getTaxSection(layout string) (taxSectionLayout string) {
	// get payment section
	reTax := regexp.MustCompile(`(?s)\[TAXES\](.*?)\[`)
	matchesTax := reTax.FindStringSubmatch(layout)
	if len(matchesTax) > 1 {
		taxSectionLayout = matchesTax[1]
	}
	return
}

func handleTaxLine(taxLine string) (string, error) {
	var tax strings.Builder
	lines := strings.Split(taxLine, "\n")
	for _, line := range lines {
		lineFormat, err := GetLineFormat(line)
		if err != nil {
			return tax.String(), err
		}
		tax.WriteString(lineFormat)
	}
	return removeEmptyLines(tax.String()), nil
}
