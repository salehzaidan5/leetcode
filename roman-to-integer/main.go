package main

func romanToInt(s string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	v := 0
	for i := 0; i < len(s); i++ {
		current := romanValues[s[i]]
		if i >= len(s)-1 {
			v += current
			continue
		}

		next := romanValues[s[i+1]]
		if current < next {
			v -= current
		} else {
			v += current
		}
	}
	return v
}
