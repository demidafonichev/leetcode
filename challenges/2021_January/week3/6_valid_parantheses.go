func opposite(br rune) rune {
	switch br {
	case '(':
		return ')'
	case '{':
		return '}'
	case '[':
		return ']'
	}
	return ' '
}

func isValid(s string) bool {
	n := len(s)
	idx, cl := -1, make([]rune, n)

	for _, br := range s {
		switch br {
		case '(', '{', '[':
			idx++
			cl[idx] = opposite(br)
		case ')', '}', ']':
			if idx != -1 && cl[idx] == br {
				idx--
			} else {
				return false
			}
		}
	}

	return idx == -1
}
