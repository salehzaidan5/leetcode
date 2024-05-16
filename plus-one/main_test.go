package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	digits []int
	want   []int
}{
	{[]int{1, 2, 3}, []int{1, 2, 4}},
	{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	{[]int{9}, []int{1, 0}},
	{[]int{9, 8, 9}, []int{9, 9, 0}},
}

func TestAdd(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.digits)
		t.Run(testname, func(t *testing.T) {
			ans := plusOne(tt.digits)
			assert.Equal(t, tt.want, ans)
		})
	}
}
