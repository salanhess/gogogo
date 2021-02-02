package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 1, 1, 2, 3, 5, 8, 13 ...
//    a, b
//       a, a+b
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//abstrat func() int to struct,then implement it's Read method
//so that printFileContents can print it out
type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 100 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum = sum + v
		return sum
	}
}
func main() {
	f := fibonacci()
	printFileContents(f)
	f2 := adder()
	for i := 0; i <= 10; i++ {
		fmt.Printf("1 + 2 + ... + %d = %d\n", i, f2(i))
	}
}
