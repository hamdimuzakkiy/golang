package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	testingValue := map[string]string{
		"hamdi ahmadi":          "idamha idmah",
		"muzakkiy":              "yikkazum",
		"test":                  "tset",
		"lol":                   "lol",
		"qwertyuiop":            "poiuytrewq",
		"325hjk32n4jkn32r92384": "48329r23nkj4n23kjh523",
		"":       "",
		"...../": "/.....",
	}

	for key, val := range testingValue {
		assert.Equal(t, val, ReverseString(key), "ReverseString")
	}
}

func TestPlus(t *testing.T) {
	assert.Equal(t, 4, Plus(1, 3), "ReverseString")
}

func TestPower(t *testing.T) {
	type pair struct {
		a int
		b int
	}

	testingValue := map[*pair]int{
		&pair{
			3, 3,
		}: 27,
		&pair{
			10, 3,
		}: 1000,
		&pair{
			2, 8,
		}: 256,
		&pair{
			9, 3,
		}: 729,
		&pair{
			1, 2,
		}: 1,
		&pair{
			101, 1,
		}: 101,
		&pair{
			100, 0,
		}: 1,
	}

	for key, val := range testingValue {
		assert.Equal(t, val, Power(key.a, key.b), "ReverseString")
	}
}

func TestSub(t *testing.T) {
	assert.Equal(t, 10, Sub(20, 10), "Tes Sub should return a -b")
}
