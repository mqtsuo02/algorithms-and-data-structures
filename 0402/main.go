package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stack : input numbers
type Stack []int

// Push : push to stack
func (sp *Stack) Push(v int) {
	*sp = append(*sp, v)
}

// Pop : pop from stack
func (sp *Stack) Pop() (v int) {
	v = (*sp)[len(*sp)-1]
	*sp = (*sp)[:len(*sp)-1]
	return
}

func main() {
	s := Stack{}
	sc := bufio.NewScanner(os.Stdin)
	if !sc.Scan() {
		return
	}
	inputs := strings.Fields(sc.Text())
	for _, input := range inputs {
		n, e := strconv.Atoi(input)
		if e == nil {
			s.Push(n)
			continue
		}
		switch input {
		case "+":
			s.Push(s.Pop() + s.Pop())
		case "-":
			s.Push(s.Pop()*-1 + s.Pop())
		case "*":
			s.Push(s.Pop() * s.Pop())
		}
	}
	fmt.Println(s.Pop())
}
