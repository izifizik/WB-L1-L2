package main

import (
	"Work/L2/pattern/facade/model"
	"fmt"
)

/*
+: изоляция
*/
func main() {
	tBank := model.Bank{
		Name:  "T-bank",
		Cards: []model.Card{},
	}
	card1 := model.Card{
		Name:    "T-card-1",
		Bank:    &tBank,
		Balance: 0,
	}
	card2 := model.Card{
		Name:    "T-card-2",
		Bank:    &tBank,
		Balance: 1000 - 7,
	}
	user1 := model.User{
		Name: "Andrey",
		Card: &card1,
	}
	user2 := model.User{
		Name: "ZXC-gul'",
		Card: &card2,
	}
	cheese := model.Product{
		Name:  "СЫР",
		Price: 110,
	}
	bottle := model.Product{
		Name:  "Bottle",
		Price: 1000,
	}
	shop := model.Shop{
		Name: "WB",
		Product: []model.Product{
			cheese,
			bottle,
		},
	}

	tBank.Cards = append(tBank.Cards, card1, card2)

	if !user1.Buy(shop, "СЫР") {
		fmt.Println("Нет сыра \\ денег")
	} else {
		fmt.Println("user 1 successfully buy CHEESE")
	}

	if !user2.Buy(shop, "Bottle") {
		fmt.Println("Нет ботела \\ денег")
		return
	}
	fmt.Println("user 2 successfully buy ботел")

}
