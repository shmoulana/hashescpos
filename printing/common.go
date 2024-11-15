package printing

import (
	"errors"
	"fmt"
	"strings"
)

const SizeOfReceipt = 52

const (
	Dash       = "----------------------------------------------------"
	DoubleDash = "===================================================="
	Space      = "                                                    "
)

const (
	AlineCenter  = "Center"
	AlineJustify = "Justify"
)

// const (
// 	TagTitle   = "Title"
// 	TagFull    = "Full"
// 	TagCenter  = "Center"
// 	TagJustify = "Justify"
// )

const (
	TagT = "T"
	TagF = "F"
	TagC = "C"
	TagJ = "J"
)

// var mapStringToTag = map[string]string{
// 	TagTitle:   TagT,
// 	TagFull:    TagF,
// 	TagCenter:  TagC,
// 	TagJustify: TagJ,
// }

func centerlizeReceipt(s string) string {
	if s == "" {
		return "\n"
	}
	padding := (SizeOfReceipt - len(s)) / 2
	return fmt.Sprintf("%*s%s%*s\n", padding, "", s, padding, "")
}

func justifyRecept(s string) string {
	if s == "" {
		return "\n"
	}
	parts := strings.Split(s, "|")
	if len(parts) == 0 {
		return s + "\n"
	}
	if len(parts) == 1 {
		return parts[0] + "\n"
	}
	left, right := parts[0], parts[1]
	ans := Space[len(left):]
	ans = left + ans
	ans = ans[:len(ans)-len(right)]
	ans = ans + right
	ans = ans + "\n"
	return ans
}

func validateLine(line string) error {
	if len(line) > SizeOfReceipt {
		return errors.New("line to long, line: " + line)
	}
	return nil
}

var LayoutTest = `
[LAYOUT]
<EB>
<C10>CURRY VILLAGE
<C10>BANANA LEAF P/L
<DB>
<C>8 LIM TECK KIM ROAD
<C>TEL : 6226 2562
<F>-
<C10>Receipt No: 1
<J00>Date: | 01-02-2006 15:12:13
{ENTITIES}
<J00> Qty Items|Price  Amount
{ORDERS}
<EB>
{DISCOUNTS}
<J10>Total:|{TICKET TOTAL}
<J10>{TICKET TAG:PersonCount}|
{PAYMENTS}
<DB>
<F>=
<C10>THANK YOU
[DISCOUNTS]
<J00>Name % AMOUNT|TOTAL
[PAYMENTS]
<J00>Cash | SomeThing
[ORDERS]
<J00>   1 Highland Coffee |  15.00 15.00
<J00>   1 Red Bull |  18.50 18.50
{ORDER TAGS}
[ORDERS:Gift]
<J00>- 1 Kitty|**GIFT**
<J00>- 2 Flower|**GIFT**
{ORDER TAGS}
[ORDER TAGS]
<J00> * Tag Name 1 | Tag Price 1
<J00> * Tag Name 2 | Tag Price 2
[ENTITIES:Table]
<L00>Table: B14
[ENTITIES:Customer]
<J00>Customer: Julie Tran | 0393698457
`
