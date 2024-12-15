package main

import (
	"printing-sampa-pos/model"
	"printing-sampa-pos/pos"
	"time"
)

func main() {
	templatePath := "data/template/template.txt"
	// Please keep in mind that option layout have not supported yet
	layoutPath := "data/layout/layout.txt"
	// Init pos handler with template path
	posHandlerTemplate, err := pos.NewPosFromTemplateFile(templatePath, layoutPath)
	if err != nil {
		panic(err)
	}
	// Printing
	// assume host is localhost and port 1234
	posHandlerTemplate.PrintWithInternet("localhost", "1234")
	// ----------------------------------------------------------------------------------------
	// Printing with ticket object data
	ticket := model.Ticket{
		Index:       "1",
		Terminal:    "CashOut",
		LoginUser:   "Chronnie Do",
		PaymentDate: time.Now().Format(model.TimeFormatDate),
		PaymentTime: time.Now().Format(model.TimeFormatTime),
		PaymentType: "",
		Total:       "15.00",
		Tag: model.Tag{
			Pax:     "10",
			PaxTime: time.Now().Format(model.TimeFormatDateTime),
		},
		Payments: []model.Payment{
			{
				Name:     "Cash",
				Tendered: "Cash",
				PaymentInformation: model.PaymentInfo{
					RefNo:   "1",
					RefTime: time.Now().Format(model.TimeFormatDateTime),
				},
				Amount: "7.50",
			},
			{
				Name:     "Credit",
				Tendered: "Credit",
				PaymentInformation: model.PaymentInfo{
					RefNo:   "2",
					RefTime: time.Now().Format(model.TimeFormatDateTime),
				},
				Amount: "7.50",
			},
		},
		Orders: []model.Order{
			{
				Name:     "Red Bull",
				Quantity: "1",
				Price:    "7.50",
			},
			{
				Name:     "Coffee",
				Quantity: "1",
				Price:    "7.50",
			},
		},
		Discounts: []model.Discount{
			{
				Name:        "Halloween",
				Description: "Any thing?",
				Total:       "5.00",
			},
		},
		Taxes: []model.Tax{
			{
				Name:   "Tax",
				Rate:   "30%",
				Amount: "5.00",
			},
		},
		Entity: model.Entity{
			Table: model.EntityTable{
				Name: "Table B14",
			},
			Member: model.EntityMember{
				Name: "Julie Tran",
				Data: model.EntityData{
					Name:    "??",
					Address: "??",
				},
			},
		},
	}
	// Init Handler with ticket object
	posHandler := pos.NewPos(ticket, layoutPath)
	// Printing
	// assume host is localhost and port 1234
	posHandler.PrintWithInternet("localhost", "1234")
}
