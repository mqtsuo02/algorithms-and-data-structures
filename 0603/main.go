package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func main() {
	var n int
	fmt.Scan(&n)
	s, e := point{0, 0}, point{100, 0}
	printPoint(s)
	koch(n, s, e)
	printPoint(e)
}

func koch(n int, s, e point) {
	if n < 1 {
		return
	}
	ml := point{(2*s.x + e.x) / 3, (2*s.y + e.y) / 3}
	mr := point{(s.x + 2*e.x) / 3, (s.y + 2*e.y) / 3}
	m := point{(mr.x-ml.x)*math.Cos(math.Pi/3) - (mr.y-ml.y)*math.Sin(math.Pi/3) + ml.x, (mr.x-ml.x)*math.Sin(math.Pi/3) - (mr.y-ml.y)*math.Cos(math.Pi/3) + ml.y}
	koch(n-1, s, ml)
	printPoint(ml)
	koch(n-1, ml, m)
	printPoint(m)
	koch(n-1, m, mr)
	printPoint(mr)
	koch(n-1, mr, e)
	return
}

func printPoint(p point) {
	fmt.Printf("%.8f %.8f\n", p.x, p.y)
}
