package printing

import (
	"regexp"
	"strings"
)

func HandleDiscount(layout string) (string, error) {
	var discount strings.Builder

	discountSectionLayout := getDiscountSection(layout)
	discountSection, err := handleDiscountLine(discountSectionLayout)
	if err != nil {
		return "", err
	}
	discount.WriteString(discountSection)
	return discount.String(), nil
}

func getDiscountSection(layout string) (discountSectionLayout string) {
	// get order section
	reDiscounts := regexp.MustCompile(`(?s)\[DISCOUNTS\](.*?)\[`)
	matchesDiscounts := reDiscounts.FindStringSubmatch(layout)
	if len(matchesDiscounts) > 1 {
		discountSectionLayout = matchesDiscounts[1]
	}

	return
}

func handleDiscountLine(layoutLine string) (string, error) {
	var discount strings.Builder
	lines := strings.Split(layoutLine, "\n")
	for _, line := range lines {
		lineFormat, err := GetLineFormat(line)
		if err != nil {
			return discount.String(), err
		}
		discount.WriteString(lineFormat)
	}
	return removeEmptyLines(discount.String()), nil
}
