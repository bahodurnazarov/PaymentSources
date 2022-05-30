package pkg

import (
	"paymentSource/types"
)

const depositLimit = 50_000_00

func Deposit(card types.Card, amount types.Money) types.Card {

	if amount < 0 {
		return card
	}
	if amount > depositLimit {
		return card
	}
	if !card.Active  {
		return card
	}
	card.Balance += amount
	
	return card
	
	}
	