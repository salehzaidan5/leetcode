package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	list1, list2 *ListNode
	want         *ListNode
}{
	{NewList(1, 2, 4), NewList(1, 3, 4), NewList(1, 1, 2, 3, 4, 4)},
	{NewList(), NewList(), NewList()},
	{NewList(), NewList(0), NewList(0)},
}

func TestMergeTwoSortedLists(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.list1, tt.list2)
		t.Run(testname, func(t *testing.T) {
			ans := mergeTwoLists(tt.list1, tt.list2)
			assert.Equal(t, tt.want.String(), ans.String())
		})
	}
}
