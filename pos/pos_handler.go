package pos

import (
	"errors"

	"hashescpos/model"
	"hashescpos/parsing"
	"hashescpos/utils"
)

type Pos struct {
	ticket     model.Ticket
	layoutPath string
}

func NewPos(ticket model.Ticket, layoutPath string) *Pos {
	return &Pos{
		ticket:     ticket,
		layoutPath: layoutPath,
	}
}

func NewPosFromTemplateFile(templatePath, layoutPath string) (*Pos, error) {
	if templatePath == "" {
		return nil, errors.New("missing template path")
	}

	if layoutPath == "" {
		return nil, errors.New("missing layout path")
	}

	template, err := utils.ReadFromFile(templatePath)
	if err != nil {
		return nil, err
	}

	ticket := parsing.FromTemplateToTicket(template)
	return NewPos(ticket, layoutPath), nil
}
