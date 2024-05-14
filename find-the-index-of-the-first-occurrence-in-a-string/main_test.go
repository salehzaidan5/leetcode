package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	haystack, needle string
	want             int
}{
	{"sadbutsad", "sad", 0},
	{"leetcode", "leeto", -1},
}

func TestFindIndexFirstOccurence(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.haystack, tt.needle)
		t.Run(testname, func(t *testing.T) {
			ans := strStr(tt.haystack, tt.needle)
			assert.Equal(t, tt.want, ans)
		})
	}
}
