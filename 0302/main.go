package main

import "fmt"

func main() {
	n, a := 0, [100]int{}

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	trace(&a, n)
	insertionSort(&a, n)
}

func trace(a *[100]int, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", a[i])
	}
	fmt.Printf("\n")
}

func insertionSort(a *[100]int, n int) {
	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for j >= 0 && a[j] > v {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = v
		trace(a, n)
	}
}
