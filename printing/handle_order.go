package printing

import (
	"regexp"
	"strings"
)

func HandleOrder(layout string) (string, error) {
	var order strings.Builder

	ordersSectionLayout, ordersGiftSectionLayout, ordersTagSectionLayout, ordersPromotionSectionLayout, ordersCompSectionLayout := getOrdersSections(layout)
	ordersSectionLayout = strings.ReplaceAll(ordersSectionLayout, "{ORDER TAGS}", ordersTagSectionLayout)
	ordersGiftSectionLayout = strings.ReplaceAll(ordersGiftSectionLayout, "{ORDER TAGS}", ordersTagSectionLayout)
	ordersSectionLayout = strings.ReplaceAll(ordersSectionLayout, "{ORDER PROMOTIONS}", ordersPromotionSectionLayout)
	orderSection, err := handleOrderLine(ordersSectionLayout)
	if err != nil {
		return order.String(), err
	}
	orderGiftSection, err := handleOrderLine(ordersGiftSectionLayout)
	if err != nil {
		return order.String(), err
	}
	ordersPromotionSection, err := handleOrderLine(ordersPromotionSectionLayout)
	if err != nil {
		return order.String(), err
	}
	ordersCompSection, err := handleOrderLine(ordersCompSectionLayout)
	if err != nil {
		return order.String(), err
	}

	order.WriteString(orderSection)
	order.WriteString("\n")
	order.WriteString(orderGiftSection)
	order.WriteString("\n")
	order.WriteString(ordersPromotionSection)
	order.WriteString("\n")
	order.WriteString(ordersCompSection)

	return order.String(), nil
}

func handleOrderLine(orderLine string) (string, error) {
	var order strings.Builder
	lines := strings.Split(orderLine, "\n")
	for _, line := range lines {
		lineFormat, err := GetLineFormat(line)
		if err != nil {
			return order.String(), err
		}
		order.WriteString(lineFormat)
	}
	return removeEmptyLines(order.String()), nil
}

func getOrdersSections(layout string) (
	ordersSectionLayout string,
	ordersGiftSectionLayout string,
	ordersTagSectionLayout string,
	ordersPromotionSectionLayout string,
	ordersCompSectionLayout string,
) {
	// get order section
	reOrders := regexp.MustCompile(`(?s)\[ORDERS\](.*?)\[`)
	matchesOrders := reOrders.FindStringSubmatch(layout)
	if len(matchesOrders) > 1 {
		ordersSectionLayout = matchesOrders[1]
	}

	// get order gifts section
	reOrdersGifts := regexp.MustCompile(`(?s)\[ORDERS:Gift\](.*?)\[`)
	matchesOrdersGifts := reOrdersGifts.FindStringSubmatch(layout)
	if len(matchesOrdersGifts) > 1 {
		ordersGiftSectionLayout = matchesOrdersGifts[1]
	}

	// get order tag section
	reOrdersTag := regexp.MustCompile(`(?s)\[ORDER TAGS\](.*?)\[`)
	matchesOrdersTag := reOrdersTag.FindStringSubmatch(layout)
	if len(matchesOrdersTag) > 1 {
		ordersTagSectionLayout = matchesOrdersTag[1]
	}

	// get order promotion section
	reOrderPromotion := regexp.MustCompile(`(?s)\[ORDER PROMOTIONS\](.*?)\[`)
	matchesOrderPromotion := reOrderPromotion.FindStringSubmatch(layout)
	if len(matchesOrderPromotion) > 1 {
		ordersPromotionSectionLayout = matchesOrderPromotion[1]
	}

	// get order comp section
	reOrderComp := regexp.MustCompile(`(?s)\[ORDER PROMOTIONS\](.*?)\[`)
	matchesOrderComp := reOrderComp.FindStringSubmatch(layout)
	if len(matchesOrderComp) > 1 {
		ordersCompSectionLayout = matchesOrderComp[1]
	}

	return
}
