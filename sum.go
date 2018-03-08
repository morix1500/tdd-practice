package money

type Sum struct {
    augend Money
    addend Money
}

func (s Sum) Reduce(to string) Money {
    return Money{
        amount: s.augend.amount + s.addend.amount,
        currency: to,
    }
}
