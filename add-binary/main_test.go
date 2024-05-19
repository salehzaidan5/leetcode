package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	a, b string
	want string
}{
	{"11", "1", "100"},
	{"1010", "1011", "10101"},
	{"1111", "1111", "11110"},
}

func TestAdd(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := addBinary(tt.a, tt.b)
			assert.Equal(t, tt.want, ans)
		})
	}
}
