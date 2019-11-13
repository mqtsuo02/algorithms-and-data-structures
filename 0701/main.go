package main

import "fmt"

const max = 500000
const sentinel = 2000000000

var ns [max]int
var la, ra [max/2 + 2]int
var cnt int

func main() {
	var n int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ns[i])
	}
	mergeSort(0, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(ns[i])
	}
	fmt.Printf("\n%d\n", cnt)
}

func mergeSort(l, r int) {
	if l+1 >= r {
		return
	}
	m := (l + r) / 2
	mergeSort(l, m)
	mergeSort(m, r)
	merge(l, m, r)
}

func merge(l, m, r int) {
	llen := m - l
	rlen := r - m
	for i := 0; i < llen; i++ {
		la[i] = ns[l+i]
	}
	for i := 0; i < rlen; i++ {
		ra[i] = ns[m+i]
	}
	la[llen] = sentinel
	ra[rlen] = sentinel
	var i, j int
	for k := l; k < r; k++ {
		cnt++
		if la[i] > ra[j] {
			ns[k] = ra[j]
			j++
			continue
		}
		ns[k] = la[i]
		i++
	}
	return
}
