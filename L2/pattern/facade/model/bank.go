package model

type Bank struct {
	Name  string
	Cards []Card
}

//payAbility - узнает у карты ее платежеспособность
func (b Bank) payAbility(name string, price float64) bool {
	for _, card := range b.Cards {
		if card.Name != name {
			continue
		}
		if card.Balance < price {
			return false
		}
		card.Balance -= price
	}
	return true
}
