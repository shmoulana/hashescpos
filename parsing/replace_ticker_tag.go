package parsing

import (
	"github.com/DoTuanAnh2k1/parsing-template/model"
	"regexp"
	"strings"
)

func replaceTickerTag(template string, ticket model.Ticket) (string, error) {
	m, err := convertFromStringToMap(ticket.Tag)
	if err != nil {
		return "", err
	}

	prefixStr := "{Ticket.Tag!"
	suffixStr := "}"
	re := regexp.MustCompile(`\{Ticket\.Tag!([^\}]+)\}`)

	matches := re.FindAllStringSubmatch(template, -1)

	for _, match := range matches {
		key := match[1]
		value, ok := m[key]
		if !ok {
			continue
		}
		template = strings.ReplaceAll(template, prefixStr+key+suffixStr, value)
	}
	return template, nil
}
