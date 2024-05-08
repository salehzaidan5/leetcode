package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	a, b int
	want int
}{
	{1, 2, 3},
	{0, 1, 1},
	{1, 1, 2},
}

func TestAdd(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := add(tt.a, tt.b)
			assert.Equal(t, tt.want, ans)
		})
	}
}
