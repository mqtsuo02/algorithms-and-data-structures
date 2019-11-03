package main

import "fmt"

func main() {
	n, vs := 0, []int{}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		v := 0
		fmt.Scan(&v)
		vs = append(vs, v)
	}

	min, p := vs[0], vs[1]-vs[0]
	for i := 1; i < n-1; i++ {
		if vs[i] < min {
			min = vs[i]
		}
		diff := vs[i+1] - min
		if diff > p {
			p = diff
		}
	}
	fmt.Println(p)
}
