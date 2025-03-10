package entity

import "errors"

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	err := t.validateAmount()
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) validateAmount() error {
	if t.Amount > 1000 {
		return errors.New("you don't have limit for this transaction")
	}
	if t.Amount < 1 {
		return errors.New("the amount must be greater than 1")
	}
	return nil
}
