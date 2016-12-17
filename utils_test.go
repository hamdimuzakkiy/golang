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
		assert.Equal(t, ReverseString(key), val, "ReverseString")
	}
}
