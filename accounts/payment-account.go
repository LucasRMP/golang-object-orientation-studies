package accounts

import (
	"fmt"
	"object-oriented/holder"
)

type PaymentAccount struct {
	Holder  holder.Holder
	Number  int
	Agency  int
	balance float32
}

func NewPaymentAccount(owner *holder.Holder, number int, agency int) *PaymentAccount {
	return &PaymentAccount{
		Holder:  *owner,
		balance: 0,
		Number:  number,
		Agency:  agency,
	}
}

func (p PaymentAccount) GetBalance() float32 {
	return p.balance
}

func (p *PaymentAccount) Withdraw(value float32) (string, error) {
	if value <= 0 {
		return "", fmt.Errorf("value must be greater than zero")
	}

	if p.balance < value {
		return "", fmt.Errorf("insufficient balance")
	}

	p.balance -= value
	return "withdraw completed", nil
}

func (p *PaymentAccount) Deposit(value float32) (string, error) {
	if value <= 0 {
		return "", fmt.Errorf("value must be greater than zero")
	}

	p.balance += value
	return "deposit completed", nil
}

func (p *PaymentAccount) Transfer(value float32, destination *PaymentAccount) bool {
	_, err := p.Withdraw(value)
	if err != nil {
		return false
	}

	_, err = destination.Deposit(value)
	return err == nil
}
