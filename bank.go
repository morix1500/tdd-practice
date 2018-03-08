package money

func Reduce(m Expression, currency string) Money {
    return m.(Money)
}
