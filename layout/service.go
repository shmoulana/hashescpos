package layout

import (
	"strings"

	"github.com/DoTuanAnh2k1/printing-sampa-pos/model"
)

func replaceService(layout string, serviceList []model.Service) string {
	var serviceOut string
	serviceOut += "[SERVICES]\n"
	serviceSectionLayoutFull, serviceSectionLayout := getSection(layout, `(?s)\[SERVICES\](.*?)\[`)

	if len(serviceList) != 0 {
		for _, service := range serviceList {
			tmp := serviceSectionLayout
			tmp = strings.ReplaceAll(tmp, "{CALCULATION NAME}", service.Name)
			tmp = strings.ReplaceAll(tmp, "{CALCULATION TOTAL}", service.Total)
			tmp += "\n"
			serviceOut = serviceOut + tmp
		}
	}

	layout = strings.ReplaceAll(layout, serviceSectionLayoutFull, serviceOut)
	return layout
}
