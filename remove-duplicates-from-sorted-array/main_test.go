package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums []int
	want []int
}{
	{[]int{1, 1, 2}, []int{1, 2}},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.nums)
		t.Run(testname, func(t *testing.T) {
			ans := removeDuplicates(tt.nums)
			assert.Equal(t, len(tt.want), ans)
			assert.Equal(t, tt.want, tt.nums[:ans])
		})
	}
}
