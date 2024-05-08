package main

import (
	"strconv"
)

func isPalindromeString(x int) bool {
	s := strconv.Itoa(x)

	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isPalindromeReverse(x int) bool {
	if x < 0 {
		return false
	}

	xCopy := x
	xReverse := 0
	for xCopy != 0 {
		n := xCopy % 10
		xReverse = 10*xReverse + n
		xCopy /= 10
	}

	return xReverse == x
}
