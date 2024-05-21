package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	x    int
	want int
}{
	{4, 2},
	{8, 2},
	{7, 2},
	{9, 3},
	{1, 1},
}

func TestAdd(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.x)
		t.Run(testname, func(t *testing.T) {
			ans := mySqrt(tt.x)
			assert.Equal(t, tt.want, ans)
		})
	}
}
