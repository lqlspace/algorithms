package medium

import (
	"math"
)

type State = byte
const (
	Start State = iota
	Sign
	Numbers
	End
)

type Dfa struct {
	state State
	table map[State][]State
}

func CreateDfa() *Dfa {
	return &Dfa{
		state:  Start,
		table:  map[State][]State{
			Start: {Start, Sign, Numbers, End},
			Sign: {End, End, Numbers, End},
			Numbers: {End, End, Numbers, End},
			End: {End, End, End, End}},
	}
}

func (d *Dfa) getState(c byte) State {
	return d.table[d.state][d.getCol(c)]
}

func (d *Dfa) getCol(c byte) int {
	switch {
	case c == ' ':
		return 0
	case c == '+', c == '-':
		return 1
	case isDigit(c):
		return 2
	default:
		return 3
	}
}

func myAtoi(s string) int {
	df := CreateDfa()
	n := len(s)
	ans, sign := 0, 1
	for i := 0; i < n; i++ {
		df.state = df.getState(s[i])
		if df.state == Numbers {
			ans = ans * 10 + int(s[i] - '0')
			if v, ok := overflow(ans * sign); ok {
				return v
			}
		} else if df.state == Sign && s[i] == '-' {
			sign *= -1
		}
	}

	return ans *sign
}


func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func overflow(val int) (int, bool) {
	if val > math.MaxInt32 {
		return math.MaxInt32, true
	} else if val  < math.MinInt32 {
		return math.MinInt32, true
	}

	return 0, false
}
