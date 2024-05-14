package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums   []int
	target int
	want   int
}{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
	{[]int{1, 3, 5, 6}, 0, 0},
}

func TestSearchInsert(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%d", tt.nums, tt.target)
		t.Run(testname, func(t *testing.T) {
			ans := searchInsert(tt.nums, tt.target)
			assert.Equal(t, tt.want, ans)
		})
	}
}
