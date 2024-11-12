package main

import (
	"github.com/DoTuanAnh2k1/parsing-template/model"
	"github.com/DoTuanAnh2k1/parsing-template/parsing"
	"time"
)

func main() {
	dataTicket := model.Ticket{
		Terminal:    "Terminal",
		LoginUser:   "Chronical Do",
		PaymentDate: time.Now().Format("2006-01-02"),
		PaymentType: "Cash",
		Tag:         "{\"Pax\": \"100\", \"PaxTime\": \"2020/10/10\"}",
		Payments: []model.Payment{
			{
				Name:     "Payment 1",
				Tendered: "Cash",
				PaymentInformation: model.PaymentInfo{
					RefNo:   "100",
					RefTime: "2020/10/10",
				},
			},
			{
				Name:     "Payment 2",
				Tendered: "Credit",
				PaymentInformation: model.PaymentInfo{
					RefNo:   "100",
					RefTime: "2020/10/10",
				},
			},
		},
		Orders: []model.Order{
			{
				Name:     "Highland Coffee",
				Quantity: "1",
				Price:    "15",
			},
			{
				Name:     "Red Bull",
				Quantity: "1",
				Price:    "18.50",
			},
		},
	}

	inputFilePath := "data/template/template.txt"
	outputFilePath := "data/output/output.txt"
	err := parsing.ParsingFromFileToFile(inputFilePath, outputFilePath, dataTicket)
	if err != nil {
		// Handle error here
		panic(err)
	}
}
