package string

func BruteForceSearch(str, pattern string) int {
	if len(str) == 0 || len(pattern) == 0 || len(str) < len(pattern) {
		return  -1
	}

	for i := 0; i <= len(str)-len(pattern)+1; i++ {
		sub := str[i:i+len(pattern)]

	}
}
