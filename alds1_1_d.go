package answer

/*
2 <= len(n) <= 200,000
1 <= n <= 10 ** 9
*/

func maxValue(ns []int) int {
	low, diff := ns[0], ns[1]-ns[0]
	for i := 2; i < len(ns); i++ {
		diff = max(diff, ns[i]-low)
		low = min(low, ns[i])
	}
	return diff
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
