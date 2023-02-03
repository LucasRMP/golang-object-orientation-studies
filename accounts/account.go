package accounts

type verifyAccount interface {
	Withdraw(value float32) (string, error)
}

func PayInvoice(account verifyAccount, value float32) {
	account.Withdraw(value)
}
