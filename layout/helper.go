package layout

import (
	"regexp"
	"strconv"
	"strings"
)

func fromStringToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func fromStringToFloat(s string) float64 {
	value, _ := strconv.ParseFloat(s, 64)
	return value
}

func fromFloat64ToString(f float64, precision int) string {
	return strconv.FormatFloat(f, 'f', precision, 64)
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

func getSection(layout string, regexStr string) (fullSection string, valueSection string) {
	reStr := regexp.MustCompile(regexStr)
	matches := reStr.FindStringSubmatch(layout)
	if len(matches) > 1 {
		fullSection = removeLastCharacter(matches[0])
		valueSection = matches[1]
	}
	return
}

func removeLastCharacter(s string) string {
	if len(s) < 1 {
		return s
	}
	return s[:len(s)-1]
}
