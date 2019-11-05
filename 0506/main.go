package main

import "fmt"

var n, k int
var ws []int
var sum int

func main() {
	fmt.Scan(&n, &k)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		sum += v
		ws = append(ws, v)
	}
	low, high := 0, sum
	for low < high {
		mid := (low + high) / 2
		if isEnaph(mid) {
			high = mid
			continue
		}
		low = mid + 1
	}
	fmt.Println(low)
}

func isEnaph(cap int) bool {
	var tmp, m int
	for i := 0; i < n; i++ {
		if tmp+ws[i] <= cap {
			tmp += ws[i]
			continue
		}
		tmp = ws[i]
		m++
		if m > k-1 {
			return false
		}
	}
	return true
}
