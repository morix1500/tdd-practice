package doller

type Doller struct{
    Amount int
}

func (d Doller) times(cnt int) Doller {
	return Doller{
        Amount: d.Amount * cnt,
    }
}

func (d Doller) equals(object interface{}) bool {
    o := object.(Doller)
    return d.Amount == o.Amount
}
