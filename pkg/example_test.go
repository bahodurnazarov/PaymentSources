package pkg

import (
	"fmt"
	"paymentSource/types"
)

func ExampleDeposit() {

	result := Deposit(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)

	// Output: 3000000

}

func ExampleWithdraw() {

	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)

	// Output: 1000000

}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}

	fmt.Println(Total(cards))
	// Output: 3000000
}

func ExampleMax() {
	payments := []types.Payment{
		{
			ID: 4047,
			Amount: 6200,
		},
		{
			ID: 4043,
			Amount: 100,
		},
		{
			ID: 4041,
			Amount: 300,
		},
	}
	max := Max(payments)
	fmt.Println(max.Amount)

	// Output: 6200
}


