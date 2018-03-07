package money

type Franc struct{
    amount int
}

func NewFranc(amount int) *Dollar {
    franc := &Dollar{}
    franc.amount = amount
    return franc
}

func (f Franc) times(cnt int) *Franc {
	return &Franc{
        amount: f.amount * cnt,
    }
}

func (f Franc) equals(object interface{}) bool {
    o := object.(*Franc)
    return f.amount == o.amount
}
