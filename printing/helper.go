package printing

import (
	"strconv"
	"strings"
)

func convertFromStringToInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	if s == "o" || s == "O" {
		return 0, nil
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return num, nil
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
