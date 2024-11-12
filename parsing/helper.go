package parsing

import (
	"fmt"
	"github.com/DoTuanAnh2k1/parsing-template/utils"
	"strconv"
	"strings"
)

func convertStringToFloatWithPrecision(input string, decimalPlaces int) string {
	// convert from string to float64
	value, err := strconv.ParseFloat(input, 64)
	utils.CheckError(err)
	// format it
	format := fmt.Sprintf("%%.%df", decimalPlaces)
	result := fmt.Sprintf(format, value)

	return result
}

func convertStringToInt(input string) int {
	num, err := strconv.Atoi(input)
	utils.CheckError(err)
	return num
}

func removeEmptyLines(input string) string {
	lines := strings.Split(input, "\n")
	var nonEmptyLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	return strings.Join(nonEmptyLines, "\n")
}
