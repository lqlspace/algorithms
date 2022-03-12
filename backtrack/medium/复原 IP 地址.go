package medium

import (
	"strconv"
	"strings"
)

const SegCount = 4

func restoreIpAddresses(s string) []string {
	var ans []string
	segs := make([]int, SegCount)

	var dfs func(s string, segId, start int)
	dfs = func(s string, segId, start int) {
		if segId == SegCount && start == len(s) {
			var builder strings.Builder
			for i := 0; i < SegCount; i++ {
				builder.WriteString(strconv.Itoa(segs[i]))
				if i < SegCount - 1 {
					builder.WriteString(".")
				}
			}
			ans = append(ans, builder.String())
			return
		}

		if segId == SegCount || start == len(s) {
			return
		}

		if s[start] == '0' {
			segs[segId] = 0
			dfs(s, segId+1, start+1)
			return
		}

		var num int
		for end := start; end < len(s); end++ {
			num = num * 10 + int(s[end] - '0')
			if num <= 0 || num > 255 {
				break
			}
			segs[segId] = num
			dfs(s, segId+1, end+1)
		}
	}

	dfs(s, 0, 0)

	return ans
}
