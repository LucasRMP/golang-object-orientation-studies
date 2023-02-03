package accounts

import (
	"fmt"
	"object-oriented/holder"
)

type SavingsAccount struct {
	Holder    holder.Holder
	Number    int
	Agency    int
	Operation int
	balance   float32
}

func NewSavingsAccount(owner *holder.Holder, number int, agency int) *SavingsAccount {
	return &SavingsAccount{
		Holder:  *owner,
		Number:  number,
		Agency:  agency,
		balance: 0,
	}
}

func (s SavingsAccount) GetBalance() float32 {
	return s.balance
}

func (s *SavingsAccount) Withdraw(value float32) (string, error) {
	if value <= 0 {
		return "", fmt.Errorf("value must be greater than zero")
	}

	if s.balance < value {
		return "", fmt.Errorf("insufficient balance")
	}

	s.balance -= value
	return "withdraw completed", nil
}

func (s *SavingsAccount) Deposit(value float32) (string, error) {
	if value <= 0 {
		return "", fmt.Errorf("value must be greater than zero")
	}

	s.balance += value
	return "deposit completed", nil
}
