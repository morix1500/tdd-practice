package money

type Sum struct {
	augend Expression
	addend Expression
}

func (s Sum) Plus(added Expression) Expression {
	return nil
}

func (s Sum) Reduce(bank Bank, to string) Money {
	augend := s.augend.Reduce(bank, to)
	addend := s.addend.Reduce(bank, to)
	return Money{
		amount:   augend.amount + addend.amount,
		currency: to,
	}
}
