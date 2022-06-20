package types

type Money int

type Currency string

type Category string

type PAN string

type PaymentSource struct {
	Type    string // 'card'
	Number  string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
}

type Payment struct {
	ID     int
	Amount Money
	Category Category
}

type Card struct {
	Balance  Money
	Currency Currency
	Color    string
	Active   bool
}

const (
	Platinum = "Platinum"
	Gold     = "Gold"
	Silver   = "Silver"
)

const (
	Somoni  = "TJS"
	Rubls   = "RUB"
	Dollars = "USD"
)
