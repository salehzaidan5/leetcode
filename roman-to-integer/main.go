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
	i := 0
	for i < len(s)-1 {
		current := romanValues[s[i]]
		next := romanValues[s[i+1]]
		if current < next {
			v += next - current
			i++
		} else {
			v += current
		}
		i++
	}

	if i == len(s)-1 {
		v += romanValues[s[len(s)-1]]
	}

	return v
}
