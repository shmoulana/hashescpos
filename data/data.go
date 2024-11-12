package data

var Template string = `
Terminal        : {Ticket.Terminal}
Cashier         :  {LoginUser}
Date            :   {Date} {Time}
Bill            : {Ticket.PaymentDate} {Ticket.PaymentTime}
[Cover:{Ticket.Tag!Pax}]

#Ticket.Orders#
#Ticket.Discounts#
#Ticket.Serivices#
#Ticket.Taxes#
#Ticket.Payments#

##Ticket.Payments##
Tendered    :   {Payments.Name}
Change      :   {Payments.Tendered}
RefNo       :   {Payments.PaymentInformation!RefNo}

##Ticket.Orders##
Name {Orders.Name} [=FormatDecimal({Orders.Quantity},2)] [=FormatDecimal({Orders.Price},2)]`
