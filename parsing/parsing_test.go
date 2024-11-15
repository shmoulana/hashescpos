package parsing_test

import (
	"fmt"
	"testing"

	"github.com/DoTuanAnh2k1/printing-sampa-pos/parsing"

	"github.com/stretchr/testify/require"
)

func TestParsingTicketFromTemplate(t *testing.T) {
	ticket := parsing.FromTemplateToTicket(parsing.TemplateTest)
	require.NotEmpty(t, ticket)
	require.Equal(t, "Terminal", ticket.Terminal)
	require.Equal(t, "Chronical Do", ticket.LoginUser)
	require.Equal(t, "2024-11-03 16:59:34", ticket.PaymentDate)
	// require.Equal(t, "Bill", ticket.PaymentType)
	require.NotEmpty(t, ticket.Payments)
	require.NotEmpty(t, ticket.Orders)
	fmt.Println(ticket.Payments)
	fmt.Println(ticket.Orders)
	fmt.Println(ticket.Tag)
}
