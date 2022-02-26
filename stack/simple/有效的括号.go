package simple


func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}

	stack := make([]int32, 0)
	for _, v := range s {
		if val, ok := shouldPop(v); ok {
			if len(stack) == 0 || stack[len(stack)-1] != val {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, v)
	}

	return len(stack) == 0
}

func shouldPop(b int32) (int32, bool) {
	switch b {
	case ')':
		return '(', true
	case ']':
		return '[', true
	case '}':
		return '{', true
	default:
		return 0, false
	}
}
