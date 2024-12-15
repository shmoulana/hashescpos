package model

import "fmt"

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

const (
	TagTitle   = "Title"
	TagFull    = "Full"
	TagCenter  = "Center"
	TagJustify = "Justify"
)

const (
	TimeFormatDate     = "2006-01-02"
	TimeFormatTime     = "15:04:05"
	TimeFormatDateTime = "2006-01-02 15:04:05"
)

func CenterlizeReceipt(s string) string {
	padding := (SizeOfReceipt - len(s)) / 2
	return fmt.Sprintf("%*s%s%*s\n", padding, "", s, padding, "")
}

func JustifyRecept(left string, right string) string {
	ans := Space[len(left):]
	ans = left + ans
	ans = ans[:len(ans)-len(right)]
	ans = ans + right
	return ans
}
