package medium

import (
	"math"
)

type State = byte
const (
	Start State = iota
	Sign
	InNumber
	End
)

type DFA struct {
	state State
	table map[State][]State
}

func Dfa() *DFA {
	return &DFA{
		state: Start,
		table: map[State][]State{
			Start: {Start, Sign, InNumber, End},
			Sign: {End, End, InNumber, End},
			InNumber: {End, End, InNumber, End}},
	}
}

func (d *DFA) transState(c byte) State {
	return d.table[d.state][d.column(c)]
}

func (d *DFA) column(c byte) int {
	switch {
	case c == ' ':
		return 0
	case c == '+' || c == '-':
		return 1
	case isDigit(c):
		return 2
	default:
		return 3
	}
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func overflow(val int) (int, bool) {
	if val > math.MaxInt32 {
		return math.MaxInt32, true
	} else if val < math.MinInt32 {
		return math.MinInt32, true
	}

	return 0, false
}


func myAtoi(s string) int {
	dfa := Dfa()
	n := len(s)
	ans, sign := 0, 1
	for i := 0; i < n; i++ {
		dfa.state = dfa.transState(s[i])
		if dfa.state == InNumber {
			ans = ans * 10 + int(s[i]-'0')
		} else if dfa.state == Sign && s[i] == '-' {
			sign *= -1
		}else if dfa.state == End {
			return ans * sign
		}

		if v, ok := overflow(ans * sign); ok {
			return v
		}
	}

	return ans * sign
}

