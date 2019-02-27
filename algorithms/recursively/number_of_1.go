//给个整数N，求解1到N中1的个数
package recursive

func NumberOf1(n int) int {
	if n < 1 {
		return 0
	}

	count := 0
	base := 1
	round := n;

	for round > 0 {
		weight := round % 10
		round /= 10
		count += round * base
		if weight == 1 {
			count += (n % base) + 1
		} else if weight > 1 {
			count += base
		}

		base *= 10
	}

	return count
}
