package main

type stack []rune

func (s *stack) push(b rune) {
	*s = append(*s, b)
}

func (s *stack) pop() rune {
	if len(*s) == 0 {
		return -1
	}

	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func isValid(s string) bool {
	parenPair := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := make(stack, 0)
	for _, c := range s {
		for k, v := range parenPair {
			if c == k {
				stack.push(c)
				break
			}

			if c == v {
				b := stack.pop()
				if b == -1 || b != k {
					return false
				}
			}
		}
	}

	return len(stack) == 0
}
