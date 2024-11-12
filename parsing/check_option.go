package parsing

import (
	"regexp"
	"strings"
)

func checkOption(template string) string {
	re := regexp.MustCompile(`\[(.*?)\]`)

	template = re.ReplaceAllStringFunc(template, func(s string) string {
		parts := strings.Split(s, ":")
		if len(parts) < 2 || s == "[]" {
			return ""
		}
		return s
	})

	return template
}
