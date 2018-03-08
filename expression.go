package money

type Expression interface {
	Reduce(Bank, string) Money
}
