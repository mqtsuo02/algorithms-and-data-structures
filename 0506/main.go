package main

import "fmt"

var n, k int
var ws []int

func main() {
	fmt.Scan(&n, &k)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		ws = append(ws, v)
	}
	for i := 0; ; i++ {
		if isEnaph(i) {
			fmt.Println(i)
			return
		}
	}
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
