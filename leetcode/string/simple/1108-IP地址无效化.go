package simple

import (
	"strings"
)

func defangIPaddr(address string) string {
	res := new(strings.Builder)
	for _, val := range address {
		if val != '.' {
			res.WriteRune(val)
		} else {
			res.WriteString("[.]")
		}
	}

	return res.String()
}
