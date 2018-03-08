package money

type Money struct {
	amount   int
	currency string
}

func (m Money) equals(object interface{}) bool {
	o := object.(*Money)
	return m.amount == o.amount && m.currency == o.currency
}

func (m Money) times(cnt int) *Money {
	return &Money{
		amount:   m.amount * cnt,
		currency: m.currency,
	}
}
