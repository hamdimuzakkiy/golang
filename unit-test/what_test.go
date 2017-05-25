package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdontEvenKnowWhatIsThisFunctionDo(t *testing.T) {
	type pair struct {
		a int
		b int
	}

	testingValue := map[*pair]int{
		&pair{
			a: 9,
			b: 2,
		}: 1,
		&pair{
			a: 8,
			b: 5,
		}: 3,
		&pair{
			a: 101,
			b: 3,
		}: 2,
		&pair{
			a: 498219048,
			b: 31,
		}: 6,
		&pair{
			a: 124812740210,
			b: 101,
		}: 5,
		&pair{
			a: 13,
			b: 4,
		}: 1,
	}

	for key, val := range testingValue {
		assert.Equal(t, val, IdontEvenKnowWhatIsThisFunctionDo(key.a, key.b), "IdontEvenKnowWhatIsThisFunctionDo")
	}
}
