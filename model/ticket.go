package model

type Ticket struct {
	Terminal    string
	LoginUser   string
	PaymentDate string
	PaymentTime string
	PaymentType string
	Tag         string
	Payments    []Payment
	Orders      []Order
	Discounts   []Discount
	Services    []Service
	Taxes       []Tax
}

type Payment struct {
	Name               string
	Tendered           string
	PaymentInformation PaymentInfo
}

type PaymentInfo struct {
	RefNo   string
	RefTime string
}

type Order struct {
	Name     string
	Quantity string
	Price    string
}

type Discount struct {
	Name string
}

type Service struct {
	Name string
}

type Tax struct {
	Name string
}
