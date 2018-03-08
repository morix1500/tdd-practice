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
	assert.False(t, NewDollar(6).equals(NewDollar(5)), "amount is not equals")
	assert.True(t, NewFranc(5).equals(NewFranc(5)), "amount is not equals")
	assert.False(t, NewFranc(6).equals(NewFranc(5)), "amount is not equals")
}

func TestFrancMultipliction(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, five.times(2), NewFranc(10), "times function have problem")
	assert.Equal(t, five.times(3), NewFranc(15), "times function have problem")
}
