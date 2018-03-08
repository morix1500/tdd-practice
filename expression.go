package money

type Expression interface {
    Plus(Expression) Expression
	Reduce(Bank, string) Money
}
