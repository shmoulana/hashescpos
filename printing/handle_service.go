package printing

import (
	"regexp"
	"strings"
)

func HandleService(layout string) (string, error) {
	var service strings.Builder

	serviceSectionLayout := getServiceSection(layout)
	serviceSection, err := handleServiceLine(serviceSectionLayout)
	if err != nil {
		return "", err
	}
	service.WriteString(serviceSection)

	return service.String(), nil
}

func getServiceSection(layout string) (serviceSectionLayout string) {
	// get payment section
	reService := regexp.MustCompile(`(?s)\[SERVICES\](.*?)\[`)
	matchesService := reService.FindStringSubmatch(layout)
	if len(matchesService) > 1 {
		serviceSectionLayout = matchesService[1]
	}
	return
}

func handleServiceLine(serviceLine string) (string, error) {
	var service strings.Builder
	lines := strings.Split(serviceLine, "\n")
	for _, line := range lines {
		lineFormat, err := GetLineFormat(line)
		if err != nil {
			return service.String(), err
		}
		service.WriteString(lineFormat)
	}
	return removeEmptyLines(service.String()), nil
}
