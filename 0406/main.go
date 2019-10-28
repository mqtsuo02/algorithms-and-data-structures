package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stack := []int{}
	sumCache := 0
	sums := []int{}
	total := 0
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := []rune(sc.Text())
	for i, input := range inputs {
		switch input {
		case []rune(`\`)[0]:
			stack = append(stack, i)
			continue
		case []rune("/")[0]:
			if len(stack) == 0 {
				continue
			}
			v := i - stack[len(stack)-1]
			total += v
			sumCache += v
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				sums = append(sums, sumCache)
				sumCache = 0
			}
		}
	}
	if sumCache > 0 {
		sums = append(sums, sumCache)
	}
	fmt.Println(total)
	for i := 0; i < len(sums); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(sums[i])
	}
}
