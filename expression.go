package money

type Expression interface {
	Reduce(string) Money
}
