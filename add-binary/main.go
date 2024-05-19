package main

import (
	"slices"
	"strings"
)

func addBinary(a string, b string) string {
	// Swap such that a is shorter than b.
	if len(a) > len(b) {
		a, b = b, a
	}
	// Pad '0' to a until it reached the same length as b.
	var sb strings.Builder
	for i := 0; i < len(b)-len(a); i++ {
		sb.WriteByte('0')
	}
	sb.WriteString(a)
	a = sb.String()

	carry := false
	sb.Reset()
	for i := len(a) - 1; i >= 0; i-- {
		x := a[i] - '0'
		y := b[i] - '0'
		r := x + y
		if carry {
			r += 1
			carry = false
		}
		if r == 2 {
			r = 0
			carry = true
		}
		if r == 3 {
			r = 1
			carry = true
		}
		sb.WriteByte(r + '0')
	}

	if carry {
		sb.WriteByte('1')
	}

	s := []byte(sb.String())
	slices.Reverse([]byte(s))
	return string(s)
}
