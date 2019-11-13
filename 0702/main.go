package main

import "fmt"

const maxlen = 100000

var nslen int
var ns [maxlen]int

func main() {
	fmt.Scan(&nslen)
	for i := 0; i < nslen; i++ {
		fmt.Scan(&ns[i])
	}

	pi := partition(0, nslen-1)

	for i := 0; i < nslen; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		if i == pi {
			fmt.Printf("[%d]", ns[i])
			continue
		}
		fmt.Print(ns[i])
	}
	fmt.Print("\n")
}

func partition(s, e int) int {
	i := s - 1
	for j := s; j < e; j++ {
		if ns[j] <= ns[e] {
			i++
			ns[i], ns[j] = ns[j], ns[i]
		}
	}
	ns[i+1], ns[e] = ns[e], ns[i+1]
	return i + 1
}
