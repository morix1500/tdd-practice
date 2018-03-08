package money

func NewDollar(amount int) *Money {
	return &Money{
		amount: amount,
	}
}

func NewFranc(amount int) *Money {
	return &Money{
		amount: amount,
	}
}
