package main

func lengthOfLastWord(s string) int {
	j := len(s)
	for s[j-1] == ' ' {
		j--
	}

	i := j - 1
	for i > 0 && s[i] != ' ' {
		i--
	}

	if s[i] != ' ' {
		return j
	}
	return j - 1 - i
}
