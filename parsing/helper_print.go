package parsing

import (
	"strings"

	"github.com/DoTuanAnh2k1/printing-sampa-pos/utils"
)

func getTemplate(inputFileTemplatePath string) string {
	lines, _ := utils.ReadFromFileToLines(inputFileTemplatePath)

	m := make(map[string]string)
	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			m[parts[0]] = parts[1]

		}
	}
	return lines[0]
}

func getLayout(inputFileLayoutPath string) string {
	layout, _ := utils.ReadFromFile(inputFileLayoutPath)
	return layout
}
