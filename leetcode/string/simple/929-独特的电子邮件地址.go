package simple

import (
	"strings"
)

func numUniqueEmails1(emails []string) int {
	m := make(map[string]struct{})

	for _, email := range emails {
		n := len(email)
		str := new(strings.Builder)
		var plus bool

		for i := 0; i < n;{
			for i < n && (email[i] == '.' || (plus && email[i] != '@')) {
				i++
			}
			start := i
			for i < n && email[i] != '+' && email[i] != '@' && email[i] != '.' {
				i++
			}
			if email[i] == '.' {
				str.WriteString(email[start:i])
			} else if email[i] == '+' {
				str.WriteString(email[start:i])
				plus = true
			} else if email[i] == '@' {
				str.WriteString(email[start:])
				break
			}
		}

		m[str.String()] = struct{}{}
	}

	return len(m)
}


func numUniqueEmails(emails []string) int {
	m := make(map[string]struct{})

	for _, email := range emails {
		var str = &strings.Builder{}
		idx := strings.Index(email, "@")
		local := email[:idx]
		local = strings.ReplaceAll(local, ".", "")
		idxPlus := strings.Index(local, "+")
		if idxPlus > -1 {
			local = local[:idxPlus]
		}
		str.WriteString(local)
		str.WriteString(email[idx:])
		m[str.String()] = struct{}{}
	}

	return len(m)
}
