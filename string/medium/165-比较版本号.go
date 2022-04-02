package medium

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	s1, s2 := strings.Split(version1, "."), strings.Split(version2, ".")
	for i := 0; i < len(s1) || i < len(s2); i++ {
		x, y := 0, 0
		if i < len(s1) {
			x, _ = strconv.Atoi(s1[i])
		}
		if i < len(s2) {
			y, _ = strconv.Atoi(s2[i])
		}

		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
	}

	return 0
}