package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	lcp := strs[0]
	for _, s := range strs[1:] {
		for j := 0; j < min(len(lcp), len(s)); j++ {
			if lcp[j] != s[j] {
				lcp = s[:j]
				break
			}
		}

		if len(lcp) > len(s) {
			lcp = s
		}
	}
	return lcp
}
