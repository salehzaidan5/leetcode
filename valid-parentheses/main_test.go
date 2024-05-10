package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	s    string
	want bool
}{
	{"()", true},
	{"()[]{}", true},
	{"(]", false},
	{"]", false},
}

func TestValidParentheses(t *testing.T) {
	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans := isValid(tt.s)
			assert.Equal(t, tt.want, ans)
		})
	}
}
