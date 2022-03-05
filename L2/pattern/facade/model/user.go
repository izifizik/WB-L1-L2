package model

type User struct {
	Name string
	Card *Card
}

// Buy - сокрытие всей логики в 1 методе
func (u User) Buy(shop Shop, product string) bool {
	price, ok := shop.productCheck(product)
	if !ok {
		return false
	}

	if !u.payAbility(price) {
		return false
	}

	return true
}

//payAbility - хватает ли баланса пользователя?
func (u User) payAbility(price float64) bool {
	if !u.Card.payAbility(price) {
		return false
	}
	return true
}
