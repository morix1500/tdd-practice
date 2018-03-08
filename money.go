package money

type Money struct {
	amount int
}

func (m Money) equals(object interface{}) bool {
	o := object.(*Money)
	return m.amount == o.amount
}

func (m Money) times(cnt int) *Money {
	return &Money{
		amount: m.amount * cnt,
	}
}
