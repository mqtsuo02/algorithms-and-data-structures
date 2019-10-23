package main

import "fmt"

func main() {
	n, a := 0, [1000000]int{}
	m, g := 2, [100]int{4, 1}
	cnt := 0

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	shellSort(&a, n, g, m, &cnt)

	fmt.Println(m)
	for i := 0; i < m; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(g[i])
	}
	fmt.Printf("\n%d\n", cnt)
	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}
}

func shellSort(a *[1000000]int, n int, g [100]int, m int, cnt *int) {

	for i := 0; i < m; i++ {
		insertionSort(a, n, g[i], cnt)
	}
}

func insertionSort(a *[1000000]int, n int, g int, cnt *int) {
	for i := g; i < n; i++ {
		v := a[i]
		j := i - g
		for j >= 0 && a[j] > v {
			a[j+g] = a[j]
			j -= g
			*cnt++
		}
		a[j+g] = v
	}
}
