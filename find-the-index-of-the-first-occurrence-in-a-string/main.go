package main

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			break
		}

		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
