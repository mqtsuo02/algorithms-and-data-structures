package main

import "fmt"

// Card : card have suit and value
type Card struct {
	Suit  rune
	Value int
}

func main() {
	n, c := 0, [36]Card{}

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%c%d", &c[i].Suit, &c[i].Value)
	}

	bc, sc := c, c

	bubbleSort(&bc, n)
	printCards(&bc, n)
	printStability(isStable(&c, &bc, n))

	selectionSort(&sc, n)
	printCards(&sc, n)
	printStability(isStable(&c, &sc, n))
}

func bubbleSort(c *[36]Card, n int) {
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if c[j].Value < c[j-1].Value {
				c[j-1], c[j] = c[j], c[j-1]
			}
		}
	}
}

func selectionSort(c *[36]Card, n int) {
	for i := 0; i < n; i++ {
		mini := i
		for j := i; j < n; j++ {
			if c[j].Value < c[mini].Value {
				mini = j
			}
		}
		if i != mini {
			c[i], c[mini] = c[mini], c[i]
		}
	}
}

func isStable(inc, outc *[36]Card, n int) bool {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := 0; k < n; k++ {
				for l := k + 1; l < n; l++ {
					if inc[i].Value == inc[j].Value && inc[i] == outc[l] && inc[j] == outc[k] {
						return false
					}
				}
			}
		}
	}
	return true
}

func printCards(c *[36]Card, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%c%d", c[i].Suit, c[i].Value)
	}
	fmt.Print("\n")
}

func printStability(isStable bool) {
	if isStable {
		fmt.Println("Stable")
		return
	}
	fmt.Println("Not Stable")
}
