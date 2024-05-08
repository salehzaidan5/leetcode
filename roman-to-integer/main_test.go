package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	s    string
	want int
}{
	{"III", 3},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func TestRomanToInteger(t *testing.T) {
	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans := romanToInt(tt.s)
			assert.Equal(t, tt.want, ans)
		})
	}
}
