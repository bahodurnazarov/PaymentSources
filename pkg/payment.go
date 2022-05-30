package pkg

import (
	"paymentSource/types"
)

func Max(payments []types.Payment) types.Payment {

	for counter := 1; counter < len(payments); counter++ {
		if payments[0].Amount < payments[counter].Amount{
			payments[0].Amount = payments[counter].Amount
		}
	}

	return payments[0]
}
