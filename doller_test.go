package doller

import (
	"testing"
)

func TestMultipliction(t *testing.T) {
	five := Doller{Amount: 5}
	product := five.times(2)
	if got, want := product.Amount, 10; got != want {
		t.Errorf("amount: got %v want %v", got, want)
	}

    product = five.times(3)
	if got, want := product.Amount, 15; got != want {
		t.Errorf("amount: got %v want %v", got, want)
	}

}
