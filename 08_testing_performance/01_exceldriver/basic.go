package main

import (
	"fmt"
	"math"
)

func triangel() {
	var a, b int = 3, 4
	fmt.Println(calcTriangel(a, b))
}

func calcTriangel(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
