package parsing

import (
	"regexp"
	"strings"
)

func findFormatDecimal(template string) string {
	template = replaceFormatDecimalWithValueFloat(template)
	template = replaceFormatDecimalWithValueInteger(template)
	return template
}

func replaceFormatDecimalWithValueFloat(template string) string {
	// Create regex to find all format fit float
	re := regexp.MustCompile(`=FormatDecimal\([01234566789]*\.[01234566789]*\,[01234566789]*\)`)

	// Replace all formatDecimal
	return re.ReplaceAllStringFunc(template, formatDecimal)
}

func replaceFormatDecimalWithValueInteger(template string) string {
	// Create regex to find all format
	re := regexp.MustCompile(`=FormatDecimal\([01234566789]*\,[01234566789]*\)`)

	// Replace all formatDecimal
	return re.ReplaceAllStringFunc(template, formatDecimal)
}

func formatDecimal(match string) string {
	match = strings.Split(match, "(")[1]
	match = strings.Split(match, ")")[0]
	parts := strings.Split(match, ",")
	// fmt.Println("match: ", match)
	decimalPlaces := convertStringToInt(parts[1])
	value := convertStringToFloatWithPrecision(parts[0], decimalPlaces)
	return value
}
