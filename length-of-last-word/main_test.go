package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	s    string
	want int
}{
	{"Hello World", 5},
	{"   fly me   to   the moon  ", 4},
	{"luffy is still joyboy", 6},
	{"a", 1},
	{" a", 1},
}

func TestLengthOfLastWord(t *testing.T) {
	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans := lengthOfLastWord(tt.s)
			assert.Equal(t, tt.want, ans)
		})
	}
}
