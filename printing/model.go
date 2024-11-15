package printing

type Line struct {
	Format FormatLine
	IsFull bool
	Body   string
}

type FormatLine struct {
	Align       string
	FrontHeight int
	FrontWidth  int
}

type Order struct {
	Items ItemList
	Gifts GiftList
}

type ItemList struct {
	Items []OrderItem
	Tags  []OrderTag
}

type GiftList struct {
	Gifts []OrderGift
	Tags  []OrderTag
}

type OrderItem struct {
	Quantity string
	ItemName string
	Price    string
	Amount   string
}

type OrderGift struct {
	Quantity string
	GiftName string
}

type OrderTag struct {
	Name  string
	Price string
}

type Entity struct {
	Table    string
	Customer string
}
