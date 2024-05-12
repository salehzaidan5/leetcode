package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums []int
	val  int
	want []int
}{
	{[]int{3, 2, 2, 3}, 3, []int{2, 2}},
	{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 4, 0, 3}},
	{[]int{}, 0, []int{}},
	{[]int{2}, 3, []int{2}},
}

func TestAdd(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%d", tt.nums, tt.val)
		t.Run(testname, func(t *testing.T) {
			ans := removeElement(tt.nums, tt.val)
			assert.Equal(t, len(tt.want), ans)
			assert.ElementsMatch(t, tt.want, tt.nums[:ans])
		})
	}
}

func TestAddSimple(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%d", tt.nums, tt.val)
		t.Run(testname, func(t *testing.T) {
			ans := removeElementSimple(tt.nums, tt.val)
			assert.Equal(t, len(tt.want), ans)
			assert.ElementsMatch(t, tt.want, tt.nums[:ans])
		})
	}
}
