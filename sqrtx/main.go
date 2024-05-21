package main

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	i := 1
	j := x - 1
	var best int
	for i <= j {
		m := i + (j-i)/2
		if m*m > x {
			j = m - 1
		} else if m*m < x {
			best = m
			i = m + 1
		} else {
			return m
		}
	}
	return best
}
