package doller

import (
	"testing"
)

func TestMultipliction(t *testing.T) {
	five := Doller{Amount: 5}
	if got, want := five.times(2).Amount, 10; got != want {
		t.Errorf("amount: got %v want %v", got, want)
	}

	if got, want := five.times(3).Amount, 15; got != want {
		t.Errorf("amount: got %v want %v", got, want)
	}
}

func TestEquality(t *testing.T)  {
    a := Doller{Amount: 5}
    b := Doller{Amount: 5}
    if !a.equals(b) {
        t.Error("equals: got false want true")
    }
    c := Doller{Amount: 6}
    if a.equals(c) {
        t.Error("equals: got true want false")
    }
}
