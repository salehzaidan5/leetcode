package main

func lengthOfLastWord(s string) int {
	j := len(s)
	for s[j-1] == ' ' {
		j--
	}

	i := j - 1
	for i > 0 && s[i-1] != ' ' {
		i--
	}

	return j - i
}
