package money

type Pair struct {
    from string
    to   string
}

func NewPair(from string, to string) Pair {
    return Pair{
        from: from,
        to:   to,
    }
}

func (p Pair) Equals(object interface{}) bool {
    pair := object.(Pair)
    return p.from == pair.from && p.to == pair.to
}

func (p Pair) hashCode() int {
    return 0
}
