package parsing

import (
	"github.com/DoTuanAnh2k1/parsing-template/model"
	"github.com/DoTuanAnh2k1/parsing-template/utils"
	"strings"
	"time"
)

func ParsingFromFileToFile(filePathInput, filePathOutput string, ticket model.Ticket) error {
	template, err := ParsingFromFile(filePathInput, ticket)
	if err != nil {
		return err
	}
	err = utils.WriteToFile(filePathOutput, template)
	if err != nil {
		return err
	}
	return nil
}

func ParsingFromFile(filePath string, ticket model.Ticket) (string, error) {
	template, err := utils.ReadFromFile(filePath)
	utils.CheckError(err)
	return Parsing(template, ticket)
}

func ParsingToFile(filePathOutput string, template string, ticket model.Ticket) error {
	template, err := Parsing(template, ticket)
	if err != nil {
		return err
	}
	err = utils.WriteToFile(filePathOutput, template)
	if err != nil {
		return err
	}
	return nil
}

func Parsing(template string, ticket model.Ticket) (string, error) {
	timeNow := time.Now().Format("2006-01-02 15:04:23")
	parts := strings.Split(timeNow, " ")
	return ParsingWithDate(template, ticket, parts[0], parts[1])
}

func ParsingWithDate(template string, ticket model.Ticket, date, timeStr string) (string, error) {
	var err error
	template, err = replaceValue(ticket, template, date, timeStr)
	if err != nil {
		return "", err
	}
	template = findFormatDecimal(template)
	template = findFormateDate(template)
	template = checkOption(template)
	return removeEmptyLines(template), nil
}
