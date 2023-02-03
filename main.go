package main

import (
	"fmt"
	"object-oriented/accounts"
	"object-oriented/holder"
)

func main() {
	lucas := holder.NewHolder("Lucas Pessone", "123.456.789-10", "Software Engineer")

	lucasPayment := accounts.NewPaymentAccount(lucas, 123, 456)
	lucasPayment.Deposit(1000)
	accounts.PayInvoice(lucasPayment, 400)
	fmt.Println(lucasPayment)

	lucasSavings := accounts.NewSavingsAccount(lucas, 123, 456)
	lucasSavings.Deposit(500)
	accounts.PayInvoice(lucasSavings, 132.50)
	fmt.Println(lucasSavings)
}
