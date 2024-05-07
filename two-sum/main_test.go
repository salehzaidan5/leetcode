package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%d", tt.nums, tt.target)
		t.Run(testname, func(t *testing.T) {
			ans := twoSum(tt.nums, tt.target)
			assert.Equal(t, tt.want, ans)
		})
		t.Run(testname, func(t *testing.T) {
			ans := twoSumHashMap(tt.nums, tt.target)
			assert.Equal(t, tt.want, ans)
		})
	}
}
