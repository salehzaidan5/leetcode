package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.x)
		t.Run(testname, func(t *testing.T) {
			ans := isPalindromeString(tt.x)
			assert.Equal(t, tt.want, ans)
		})
	}
}
