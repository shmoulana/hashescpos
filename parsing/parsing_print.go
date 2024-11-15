package parsing

import (
	"regexp"
	"strings"

	"github.com/DoTuanAnh2k1/printing-sampa-pos/model"
	"github.com/DoTuanAnh2k1/printing-sampa-pos/utils"
)

func ParsingValue(inputFileLayoutPath, inputFileTemplatePath string) string {
	template := getTemplate(inputFileTemplatePath)
	layout := getLayout(inputFileLayoutPath)

	// return first
	return template + layout
}

func ParsingTicketFromFile(filePath string) (model.Ticket, error) {
	template, err := utils.ReadFromFile(filePath)
	if err != nil {
		return model.Ticket{}, err
	}

	return FromTemplateToTicket(template), nil
}

func FromTemplateToTicket(template string) model.Ticket {
	var ticket model.Ticket
	// Regular expressions for parsing different parts
	reTerminal := regexp.MustCompile(`(?m)^Terminal\s*:\s*(.*)`)
	reCashier := regexp.MustCompile(`(?m)^Cashier\s*:\s*(.*)`)
	reDate := regexp.MustCompile(`(?m)^Date\s*:\s*(.*)`)
	reBill := regexp.MustCompile(`(?m)^Bill\s*:\s*(.*)`)
	reTag := regexp.MustCompile(`(?m)^Cover:(\d+)`)
	rePayment := regexp.MustCompile(`(?m)^Tendered\s*:\s*(.*)\nChange\s*:\s*(.*)\nRefNo\s*:\s*(\d+)`)
	reOrder := regexp.MustCompile(`(?m)^Name\s+(.*)\s+\[(\d+\.\d+)\]\s+\[(\d+\.\d+)\]`)

	// Extracting data using regex
	if match := reTerminal.FindStringSubmatch(template); match != nil {
		ticket.Terminal = strings.TrimSpace(match[1])
	}
	if match := reCashier.FindStringSubmatch(template); match != nil {
		ticket.LoginUser = strings.TrimSpace(match[1])
	}
	if match := reDate.FindStringSubmatch(template); match != nil {
		ticket.PaymentDate = strings.TrimSpace(match[1])
	}
	if match := reBill.FindStringSubmatch(template); match != nil {
		ticket.PaymentType = strings.TrimSpace(match[1])
	}
	if match := reTag.FindStringSubmatch(template); match != nil {
		ticket.Tag.Pax = match[1]
	}

	// Extracting payments
	paymentMatches := rePayment.FindAllStringSubmatch(template, -1)
	for _, match := range paymentMatches {
		payment := model.Payment{
			Name:     match[1],
			Tendered: match[2],
			PaymentInformation: model.PaymentInfo{
				RefNo: match[3],
			},
		}
		ticket.Payments = append(ticket.Payments, payment)
	}

	// Extracting orders
	orderMatches := reOrder.FindAllStringSubmatch(template, -1)
	for _, match := range orderMatches {
		order := model.Order{
			Name:     match[1],
			Quantity: match[2],
			Price:    match[3],
		}
		ticket.Orders = append(ticket.Orders, order)
	}
	return ticket
}
