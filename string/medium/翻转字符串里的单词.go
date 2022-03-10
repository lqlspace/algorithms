package medium

import (
	"strings"
)

type WordState = byte
type WordDFA struct {
	state WordState
	ht map[WordState][]WordState
}

const (
	StartState WordState = iota
	BoundState
	InWordsState
	InterruptState
)

func CreateWordDFA() *WordDFA {
	return &WordDFA{
		state: StartState,
		ht:    map[WordState][]WordState{
			StartState: {StartState, BoundState},
			BoundState: {InterruptState, InWordsState},
			InterruptState: {StartState, BoundState},
			InWordsState: {InterruptState, InWordsState},
		},
	}
}

func (dfa *WordDFA) getWordCol(c byte) int {
	switch {
	case c == ' ':
		return 0
	default:
		return 1
	}
}

func reverseWords2(s string) string {
	dfa := CreateWordDFA()

	var ans []string
	n := len(s)
	idx := -1
	for i := n-1; i >= 0; i-- {
		dfa.state = dfa.ht[dfa.state][dfa.getWordCol(s[i])]
		switch dfa.state {
		case BoundState:
			idx = i
		case InterruptState:
			ans = append(ans, s[i+1:idx+1])
			idx = -1
		}
	}
	if idx >= 0 {
		ans = append(ans, s[:idx+1])
	}

	return strings.Join(ans, " ")
}


func reverseWords(s string) string {
	n := len(s)

	idx := -1
	var ans []string
	for i := n-1; i >= 0; i-- {
		if s[i] == ' ' {
			if idx >= 0 {
				ans = append(ans, s[i+1:idx+1])
				idx = -1
			}
			continue
		}

		if idx == -1 {
			idx = i
		}
	}

	if idx >= 0 {
		ans = append(ans, s[:idx+1])
	}

	return strings.Join(ans, " ")
}
