package main

import "fmt"

var n, m int
var na []int

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		v := 0
		fmt.Scan(&v)
		na = append(na, v)
	}
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		v := 0
		fmt.Scan(&v)
		if solve(0, v) {
			fmt.Println("yes")
			continue
		}
		fmt.Println("no")
	}
}

func solve(i, t int) bool {
	if t == 0 {
		return true
	}
	if i >= n {
		return false
	}
	return solve(i+1, t) || solve(i+1, t-na[i])
}
