package layout

import (
	"strings"

	"github.com/DoTuanAnh2k1/printing-sampa-pos/model"
)

func replaceEntityTable(layout string, entity model.EntityTable) string {
	var entityOut string
	entityOut += "[ENTITIES:Table]\n"
	entityTableSectionLayoutFull, entityTableSectionLayout := getSection(layout, `(?s)\[ENTITIES:Table\](.*?)\[`)
	entityTableSectionLayout = strings.ReplaceAll(entityTableSectionLayout, "{ENTITY NAME}", entity.Name)
	if entityTableSectionLayout != "" {
		entityOut += entityTableSectionLayout
		entityOut += "\n"
	}

	layout = strings.ReplaceAll(layout, entityTableSectionLayoutFull, entityOut)
	return layout
}

func replaceEntityMember(layout string, entity model.EntityMember) string {
	var entityOut string
	entityOut += "[ENTITIES:Member]\n"
	entityMemberSectionLayoutFull, entityMemberSectionLayout := getSection(layout, `(?s)\[ENTITIES:Member\](.*?)$`)
	entityMemberSectionLayout = strings.ReplaceAll(entityMemberSectionLayout, "{ENTITY NAME}", entity.Name)
	entityMemberSectionLayout = strings.ReplaceAll(entityMemberSectionLayout, "{ENTITY DATA Name}", entity.Data.Address)
	entityMemberSectionLayout = strings.ReplaceAll(entityMemberSectionLayout, "{ENTITY DATA Address}", entity.Data.Address)
	if entityMemberSectionLayout != "" {
		entityOut += entityMemberSectionLayout
		entityOut += "\n"
	}

	layout = strings.ReplaceAll(layout, entityMemberSectionLayoutFull, entityOut)
	return layout
}
