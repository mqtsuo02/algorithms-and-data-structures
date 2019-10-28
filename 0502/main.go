package main

import "fmt"

func main() {
	ts, tsLen := []int{}, 0
	count := 0

	fmt.Scan(&tsLen)
	for i := 0; i < tsLen; i++ {
		v := 0
		fmt.Scan(&v)
		ts = append(ts, v)
	}

	ksLen := 0
	fmt.Scan(&ksLen)
	for i := 0; i < ksLen; i++ {
		k := 0
		fmt.Scan(&k)
		if search(ts, k) {
			count++
		}
	}
	fmt.Print(count)
}

func search(ts []int, k int) bool {
	ts = append(ts, k)
	i := 0
	for ts[i] != k {
		i++
	}
	return i != len(ts)-1
}
