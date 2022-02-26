package simple

import (
	"testing"
)

func Test_isValid(t *testing.T)  {
	t.Log(isValid("()"))
	t.Log(isValid("()[]{}"))
	t.Log(isValid("([)]"))
	t.Log(isValid("[()]"))
	t.Log(isValid("[()"))
	t.Log(isValid("(("))
}
