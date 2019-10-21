package main

import "fmt"

func main() {
	n, a := 0, [100]int{}

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sw := selectionSort(&a, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", a[i])
	}
	fmt.Printf("\n%d\n", sw)
}

func selectionSort(a *[100]int, n int) int {
	sw := 0
	for i := 0; i < n; i++ {
		mini := i
		for j := i; j < n; j++ {
			if a[j] < a[mini] {
				mini = j
			}
		}
		if i != mini {
			a[i], a[mini] = a[mini], a[i]
			sw++
		}
	}
	return sw
}
