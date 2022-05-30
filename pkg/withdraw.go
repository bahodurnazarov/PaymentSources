package pkg

import (
	"paymentSource/types"
)

const withdrawLimit = 20_000_00

func Withdraw(card types.Card, amount types.Money) types.Card {

	if amount < 0 {
		return card
	}
	if amount > withdrawLimit {
		return card
	}
	if !card.Active  {
		return card
	}
	if card.Balance < amount {
		return card
	}
	card.Balance -= amount
	
	return card
	
	}
	