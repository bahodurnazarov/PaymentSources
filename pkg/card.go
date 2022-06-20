package pkg

import (
	"paymentSource/types"
)


func IssueCard(balance types.Money, currency types.Currency, color string, active bool) types.Card {

	card := types.Card{
		Balance:  balance,
		Currency: currency,
		Color:    color,
		Active:   active,
	}
	return card
}

func Total(cards []types.Card) types.Money {

	sum := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}

		sum += card.Balance
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	
	var paymentSources []types.PaymentSource
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}

		paymentSource := types.PaymentSource{
			Type:    "card",
			Balance: card.Balance,
		}
		paymentSources = append(paymentSources, paymentSource)
	}
	return paymentSources
	
	
}


func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	for _, payment := range payments {
		sum += payment.Amount
	}
	return sum / types.Money(len(payments))
}
