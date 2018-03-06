package doller

type Dollar struct{
    amount int
}

func NewDollar(amount int) *Dollar {
    dollar := &Dollar{}
    dollar.amount = amount
    return dollar
}

func (d Dollar) times(cnt int) *Dollar {
	return &Dollar{
        amount: d.amount * cnt,
    }
}

func (d Dollar) equals(object interface{}) bool {
    o := object.(*Dollar)
    return d.amount == o.amount
}
