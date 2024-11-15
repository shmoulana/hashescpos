package layout

import (
	"strings"

	"github.com/DoTuanAnh2k1/printing-sampa-pos/model"
)

func replaceOrder(layout string, orderList []model.Order) string {
	var orderOut string
	ordersSectionLayoutFull, ordersSectionLayout := getSection(layout, `(?s)\[ORDERS\](.*?)\[`)
	orderOut = orderOut + "[ORDERS]\n"
	if len(orderList) != 0 {
		for _, order := range orderList {
			quantity := fromStringToInt(order.Quantity)
			price := fromStringToFloat(order.Price)
			totalPrice := float64(quantity) * price
			tmp := ordersSectionLayout
			tmp = strings.ReplaceAll(tmp, "{NAME}", order.Name)
			tmp = strings.ReplaceAll(tmp, "{QUANTITY}", order.Quantity)
			tmp = strings.ReplaceAll(tmp, "{PRICE}", order.Price)
			tmp = strings.ReplaceAll(tmp, "{TOTAL}", fromFloat64ToString(totalPrice, 2))
			tmp += "\n"
			orderOut = orderOut + tmp
		}
	}

	layout = strings.ReplaceAll(layout, ordersSectionLayoutFull, orderOut)

	return layout
}

func replaceOrderTag(layout string, orderTagList []model.OrderTag) string {
	var orderTagOut string
	orderTagOut = orderTagOut + "[ORDERS TAGS]\n"
	ordersTagSectionLayoutFull, ordersTagSectionLayout := getSection(layout, `(?s)\[ORDER TAGS\](.*?)\[`)

	if len(orderTagList) != 0 {
		for _, tag := range orderTagList {
			tmp := ordersTagSectionLayout
			tmp = strings.ReplaceAll(tmp, "{ORDER TAG NAME}", tag.Name)
			tmp = strings.ReplaceAll(tmp, "{ORDER TAG PRICE}", tag.Price)
			tmp += "\n"
			orderTagOut = orderTagOut + tmp
		}
	}

	layout = strings.ReplaceAll(layout, ordersTagSectionLayoutFull, orderTagOut)

	return layout
}

func replaceOrderComp(layout string, orderCompList []model.OrderComp) string {
	var orderCompOut string
	ordersCompSectionLayoutFull, ordersCompSectionLayout := getSection(layout, `(?s)\[ORDERS:Comp\](.*?)\[`)
	orderCompOut = orderCompOut + "[ORDERS:Comp]\n"

	if len(orderCompList) != 0 {
		for _, comp := range orderCompList {
			tmp := ordersCompSectionLayout
			tmp = strings.ReplaceAll(tmp, "{NAME}", comp.Name)
			tmp = strings.ReplaceAll(tmp, "{QUANTITY}", comp.Quantity)
			tmp = strings.ReplaceAll(tmp, "{PRICE}", comp.Price)
			tmp = strings.ReplaceAll(tmp, "{TOTAL}", comp.Total)
			tmp += "\n"
			orderCompOut = orderCompOut + tmp
		}
	}

	layout = strings.ReplaceAll(layout, ordersCompSectionLayoutFull, orderCompOut)

	return layout
}

func replaceOrderPromotion(layout string, orderPromotionList []model.OrderPromotion) string {
	var orderPromotionOut string
	orderPromotionOut = orderPromotionOut + "[ORDERS PROMOTIONS]\n"
	orderPromotionSectionLayoutFull, orderPromotionSectionLayout := getSection(layout, `(?s)\[ORDER PROMOTIONS\](.*?)\[`)

	if len(orderPromotionList) != 0 {
		for _, comp := range orderPromotionList {
			tmp := orderPromotionSectionLayout
			tmp = strings.ReplaceAll(tmp, "{ORDER PROMOTION NAME}", comp.Name)
			tmp = strings.ReplaceAll(tmp, "{ORDER PROMOTION TOTAL}", comp.Total)
			tmp += "\n"
			orderPromotionOut = orderPromotionOut + tmp
		}
	}

	layout = strings.ReplaceAll(layout, orderPromotionSectionLayoutFull, orderPromotionOut)
	return layout
}
