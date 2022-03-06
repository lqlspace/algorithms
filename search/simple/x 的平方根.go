package simple


// 二分查找，递归
func mySqrt(x int) int {
	return binarySqrt(x, 0, x)
}

func binarySqrt(x, low, high int) int {
	if low > high {
		return high
	}

	mid := low + (high-low) >> 1
	if mid * mid == x {
		return mid
	} else if mid * mid > x {
		high = mid-1
	} else {
		low = mid + 1
	}

	return binarySqrt(x, low, high)
}

func mySqrt2(x int) int {
	low, high := 0, x
	ans := -1
	for low <= high {
		mid := low + (high-low) >> 1
		if mid * mid <= x {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return ans
}
