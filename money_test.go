package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultipliction(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, five.times(2), NewDollar(10), "times function have problem")
	assert.Equal(t, five.times(3), NewDollar(15), "times function have problem")
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(5).equals(NewDollar(5)), "amount is not equals")
	assert.False(t, NewDollar(6).equals(NewDollar(5)), "amount is equals")
	assert.False(t, NewDollar(5).equals(NewFranc(5)), "amount is equals")
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(5).currency, "currency is not USD")
	assert.Equal(t, "CHF", NewFranc(5).currency, "currency is not CHF")
}
