package parentheses

import (
	"parentheses_validation/pkg/stack"
)

func IsValid(s string) bool {
	return isValid(s)
}

func isValid(s string) bool {
	var (
		st             = stack.New[rune]()
		openingBracket rune
	)

	for _, r := range s {
		switch r {

		case '{', '[', '(': // , '<'
			st.Push(r)

		case '}', ']', ')': // , '>'
			if st.IsEmpty() {
				return false
			}

			switch r {
			case ')':
				openingBracket = r - 1
			default:
				openingBracket = r - 2
			}

			if openingBracket != st.Pop() {
				return false
			}

		default:
			continue
		}
	}

	return st.IsEmpty()
}
