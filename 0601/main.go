package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input []int
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	vs := strings.Fields(sc.Text())
	for i := 0; i < len(vs); i++ {
		v, _ := strconv.Atoi(vs[i])
		input = append(input, v)
	}
	fmt.Println(findMaximum(input, 0, len(input)))
}

func findMaximum(input []int, l, r int) int {
	if l == r-1 {
		return input[l]
	}
	m := (l + r) / 2
	lm := findMaximum(input, l, m)
	rm := findMaximum(input, m, r)
	return max(lm, rm)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
