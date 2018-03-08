package money

func Reduce(source Expression, to string) Money {
	return source.Reduce(to)
}
