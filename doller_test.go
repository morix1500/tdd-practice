package doller

import (
	"testing"
)

func TestMultipliction(t *testing.T) {
	five := NewDollar(5)
    want := NewDollar(10)
	if !five.times(2).equals(want) {
		t.Errorf("amount: got %v want %v", five.times(2), want)
	}

    want = NewDollar(15)
	if !five.times(3).equals(want) {
		t.Errorf("amount: got %v want %v", five.times(3), want)
	}
}

func TestEquality(t *testing.T)  {
    a := NewDollar(5)
    b := NewDollar(5)
    if !a.equals(b) {
        t.Error("equals: got false want true")
    }
    c := NewDollar(6)
    if a.equals(c) {
        t.Error("equals: got true want false")
    }
}
