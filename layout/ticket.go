package layout

import (
	"strings"

	"github.com/DoTuanAnh2k1/printing-sampa-pos/model"
)

func replaceTicket(layout string, ticket model.Ticket) string {
	layout = strings.ReplaceAll(layout, "{TICKET NO}", ticket.Index)
	layout = strings.ReplaceAll(layout, "{CURRENT TERMINAL}", ticket.Terminal)
	layout = strings.ReplaceAll(layout, "{Ticket.LoginUser}", ticket.LoginUser)
	layout = strings.ReplaceAll(layout, "{TICKET PAYMENT DATE}", ticket.PaymentDate)
	layout = strings.ReplaceAll(layout, "{TICKET PAYMENT TIME}", ticket.PaymentTime)
	layout = strings.ReplaceAll(layout, "{Ticket.PaymentType}", ticket.PaymentType)
	layout = strings.ReplaceAll(layout, "{TICKET TAG:PAX}", ticket.Tag.Pax)
	layout = strings.ReplaceAll(layout, "{TICKET TAG:PAXTIME}", ticket.Tag.PaxTime)
	layout = strings.ReplaceAll(layout, "{TICKET TOTAL}", ticket.Total)

	return layout
}
