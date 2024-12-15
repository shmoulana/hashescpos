package model

type Ticket struct {
	Index          string
	Terminal       string
	LoginUser      string
	PaymentDate    string
	PaymentTime    string
	PaymentType    string
	Total          string
	TagParsing     string
	Tag            Tag
	Payments       []Payment
	Orders         []Order
	OrderTags      []OrderTag
	OrderComps     []OrderComp
	OrderPromotion []OrderPromotion
	Discounts      []Discount
	Services       []Service
	Taxes          []Tax
	Entity         Entity
}

type Tag struct {
	Pax     string
	PaxTime string
}

type Payment struct {
	Name               string
	Tendered           string
	PaymentInformation PaymentInfo
	Amount             string
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

type OrderTag struct {
	Name  string
	Price string
}

type OrderComp struct {
	Name     string
	Quantity string
	Price    string
	Total    string
}

type OrderPromotion struct {
	Name  string
	Total string
}

type Discount struct {
	Name        string
	Description string
	Total       string
}

type Service struct {
	Name  string
	Total string
}

type Tax struct {
	Name   string
	Rate   string
	Amount string
}

type Entity struct {
	Table  EntityTable
	Member EntityMember
}

type EntityTable struct {
	Name string
}

type EntityMember struct {
	Name string
	Data EntityData
}

type EntityData struct {
	Name    string
	Address string
}
