package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	strs []string
	want string
}{
	{[]string{"flower", "flow", "flight"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
}

func TestLCP(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.strs)
		t.Run(testname, func(t *testing.T) {
			ans := longestCommonPrefix(tt.strs)
			assert.Equal(t, tt.want, ans)
		})
	}
}
