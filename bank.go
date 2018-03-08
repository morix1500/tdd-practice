package money

type Bank struct {
    rates map[Pair]int
}

func NewBank() Bank {
    return Bank{
        rates: map[Pair]int{},
    }
}

func (b Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(b, to)
}

func (b Bank) AddRate(from string, to string, rate int) {
    b.rates[NewPair(from, to)] = rate
}

func (b Bank) Rate(from string, to string) int {
    rate, exists := b.rates[NewPair(from, to)]
    if exists {
        return rate
    }
    return 1
}
