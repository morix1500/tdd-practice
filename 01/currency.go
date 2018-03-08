package money

func NewDollar(amount int) Money {
	return Money{
		amount:   amount,
		currency: "USD",
	}
}

func NewFranc(amount int) Money {
	return Money{
		amount:   amount,
		currency: "CHF",
	}
}
