package model

type Card struct {
	Name    string
	Bank    *Bank
	Balance float64
}

//GetBalance - получает остаток на карте
func (c Card) GetBalance() float64 {
	return c.Balance
}

//payAbility - хватает ли баланса карты?
func (c Card) payAbility(price float64) bool {
	return c.Bank.payAbility(c.Name, price)
}
