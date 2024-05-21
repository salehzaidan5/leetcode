package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	n    int
	want int
}{
	{2, 2},
	{3, 3},
}

func TestAdd(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.n)
		t.Run(testname, func(t *testing.T) {
			ans := climbStairs(tt.n)
			assert.Equal(t, tt.want, ans)
		})
	}
}
