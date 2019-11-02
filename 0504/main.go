package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const m int = 1046527

var hashArray [m]string

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	cs := [][]string{}
	for i := 0; i < n; i++ {
		sc.Scan()
		cs = append(cs, strings.Fields(sc.Text()))
	}
	for _, c := range cs {
		switch c[0] {
		case "insert":
			insert(c[1])
		case "find":
			if find(c[1]) {
				fmt.Println("yes")
				continue
			}
			fmt.Println("no")
		}
	}
}

func insert(v string) {
	key := getKey(v)
	for i := 0; i < m; i++ {
		hash := (makeHash1(key) + makeHash2(key)*i) / m
		if hashArray[hash] == "" {
			hashArray[hash] = v
			return
		}
	}
}

func find(v string) bool {
	key := getKey(v)
	for i := 0; i < m; i++ {
		hash := (makeHash1(key) + makeHash2(key)*i) / m
		switch hashArray[hash] {
		case "":
			return false
		case v:
			return true
		}
	}
	return false
}

func makeHash1(k int) int {
	return k % m
}

func makeHash2(k int) int {
	return k%(m-1) + 1
}

func getKey(v string) int {
	sum, p := 0, 1
	vs := strings.Split(v, "")
	for i := 0; i < len(vs); i++ {
		sum += stoi(vs[i]) * p
		p *= 5
	}
	return sum
}

func stoi(s string) int {
	switch s {
	case "A":
		return 1
	case "C":
		return 2
	case "G":
		return 3
	case "T":
		return 4
	default:
		return 0
	}
}
