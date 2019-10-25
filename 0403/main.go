package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Process : process name and time required
type Process struct {
	Name string
	Time int
}

// ProcessQueue : queue of process
type ProcessQueue []Process

// Enqueue : add process to tail of queue
func (pq *ProcessQueue) Enqueue(p Process) {
	*pq = append(*pq, p)
}

// Dequeue : take out process from head of queue
func (pq *ProcessQueue) Dequeue() (p Process, e error) {
	if len(*pq) == 0 {
		e = errors.New("there is no process in queue")
		return
	}
	p = (*pq)[0]
	*pq = (*pq)[1:]
	return
}

// IsEmpty : check if queue is empty
func (pq *ProcessQueue) IsEmpty() bool {
	if len(*pq) == 0 {
		return true
	}
	return false
}

func main() {
	pq := ProcessQueue{}
	tt := 0

	sc := bufio.NewScanner(os.Stdin)
	if !sc.Scan() {
		return
	}

	fs := strings.Fields(sc.Text())
	n, _ := strconv.Atoi(fs[0])
	q, _ := strconv.Atoi(fs[1])

	for i := 0; i < n; i++ {
		sc.Scan()
		fs = strings.Fields(sc.Text())
		t, _ := strconv.Atoi(fs[1])
		pq.Enqueue(Process{Name: fs[0], Time: t})
	}

	for !pq.IsEmpty() {
		p, _ := pq.Dequeue()
		if p.Time <= q {
			tt += p.Time
			fmt.Println(p.Name, tt)
			continue
		}
		pq.Enqueue(Process{Name: p.Name, Time: p.Time - q})
		tt += q
	}
}
