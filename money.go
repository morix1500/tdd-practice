package money

type Money struct {
	amount   int
	currency string
}

func (m Money) Equals(object interface{}) bool {
	o := object.(Money)
	return m.amount == o.amount && m.currency == o.currency
}

func (m Money) Times(cnt int) Expression {
	return Money{
		amount:   m.amount * cnt,
		currency: m.currency,
	}
}

func (m Money) Plus(addend Expression) Expression {
	return Sum{
		augend: m,
		addend: addend,
	}
}

func (m Money) Reduce(bank Bank, to string) Money {
    rate := bank.Rate(m.currency, to)
    return Money{
        amount: m.amount / rate,
        currency: to,
    }
}
