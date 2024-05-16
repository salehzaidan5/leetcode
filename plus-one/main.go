package main

import "slices"

func plusOne(digits []int) []int {
	carry := false
	done := false
	for i := len(digits) - 1; i >= 0 && !done; i-- {
		if carry {
			if digits[i] == 9 {
				digits[i] = 0
				continue
			}

			digits[i] += 1
			done = true
			carry = false
			continue
		}

		if digits[i] == 9 {
			digits[i] = 0
			carry = true
			continue
		}

		digits[i] += 1
		done = true
		break
	}

	if carry {
		digits = slices.Insert(digits, 0, 1)
	}

	return digits
}
