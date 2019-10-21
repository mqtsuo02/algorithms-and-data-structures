package main

import "fmt"

func main() {
	n, a := 0, [100]int{}

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sw := bubbleSort(&a, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", a[i])
	}
	fmt.Printf("\n%d\n", sw)
}

func bubbleSort(a *[100]int, n int) int {
	sw := 0
	flag := true
	for i := 0; flag; i++ {
		flag = false
		for j := n - 1; j > i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				sw++
				flag = true
			}
		}
	}
	return sw
}
