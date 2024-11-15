package printing

import (
	"regexp"
	"strings"
)

func HandleEntity(layout string) (string, error) {
	var entity strings.Builder

	entityTableSectionLayout, entityMemberSectionLayout := getEntitySections(layout)
	entityTableSection, err := handleEntityLine(entityTableSectionLayout)
	if err != nil {
		return entity.String(), err
	}
	entityMemberSection, err := handleEntityLine(entityMemberSectionLayout)
	if err != nil {
		return entity.String(), err
	}

	entity.WriteString(entityTableSection)
	entity.WriteString("\n")
	entity.WriteString(entityMemberSection)

	return entity.String(), nil
}

func handleEntityLine(entityLine string) (string, error) {
	var entity strings.Builder
	lines := strings.Split(entityLine, "\n")
	for _, line := range lines {
		lineFormat, err := GetLineFormat(line)
		if err != nil {
			return entity.String(), err
		}
		entity.WriteString(lineFormat)
	}
	return removeEmptyLines(entity.String()), nil
}

func getEntitySections(layout string) (entityTableSectionLayout string, entityMemberSectionLayout string) {
	// get entity table section
	reEntityTable := regexp.MustCompile(`(?s)\[ENTITIES:Table\](.*?)\[`)
	matchesEntity := reEntityTable.FindStringSubmatch(layout)
	if len(matchesEntity) > 1 {
		entityTableSectionLayout = matchesEntity[1]
	}

	// get entity member section
	reEntityMember := regexp.MustCompile(`(?s)\[ENTITIES:Member\](.*?)\[`)
	matchesentitysGifts := reEntityMember.FindStringSubmatch(layout)
	if len(matchesentitysGifts) > 1 {
		entityMemberSectionLayout = matchesentitysGifts[1]
	}

	return
}
