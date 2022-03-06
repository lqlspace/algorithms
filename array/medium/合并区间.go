package medium

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return  intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{intervals[0]}
	for i := 1; i < n; i++ {
		if ans[len(ans)-1][1] >= intervals[i][0] {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
