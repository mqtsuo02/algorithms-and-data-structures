package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	key        int
	prev, next *node
}

type command struct {
	name string
	val  int
}

var head node

func main() {
	head = node{prev: &head, next: &head}
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		vs := strings.Fields(sc.Text())
		switch vs[0] {
		case "insert":
			v, _ := strconv.Atoi(vs[1])
			insert(v)
		case "delete":
			v, _ := strconv.Atoi(vs[1])
			deleteKey(v)
		case "deleteFirst":
			deleteFirst()
		case "deleteLast":
			deleteLast()
		}
	}
	print()
}

func print() {
	cur := head.next
	for i := 0; cur != &head; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(cur.key)
		cur = cur.next
	}
	fmt.Print("\n")
}

func insert(v int) {
	newNode := node{key: v, prev: &head, next: head.next}
	head.next.prev = &newNode
	head.next = &newNode
}

func deleteKey(v int) {
	t := searchKey(v)
	delete(t)
}

func searchKey(v int) *node {
	cur := head.next
	for cur != &head {
		if cur.key == v {
			return cur
		}
		cur = cur.next
	}
	return cur
}

func deleteFirst() {
	delete(head.next)
}

func deleteLast() {
	delete(head.prev)
}

func delete(t *node) {
	if t == &head {
		return
	}
	t.prev.next = t.next
	t.next.prev = t.prev
}
