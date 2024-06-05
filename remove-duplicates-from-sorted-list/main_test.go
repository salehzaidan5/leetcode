package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	head *ListNode
	want *ListNode
}{
	{NewList(1, 1, 2), NewList(1, 2)},
	{NewList(1, 1, 2, 3, 3), NewList(1, 2, 3)},
	{NewList(1, 1, 1), NewList(1)},
}

func TestDeleteDuplicates(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.head)
		t.Run(testname, func(t *testing.T) {
			ans := deleteDuplicates(tt.head)
			assert.Equal(t, tt.want.String(), ans.String())
		})
	}
}
