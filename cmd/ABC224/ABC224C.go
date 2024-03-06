package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := os.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
}

func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nis(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	return a
}

func ni2() (int, int) {
	return ni(), ni()
}

func ni3() (int, int, int) {
	return ni(), ni(), ni()
}

func ns() string {
	sc.Scan()
	return sc.Text()
}

type IntPar [2]int

func main() {
	defer wtr.Flush()
	count := 0
	n := ni()
	sl := make([]IntPar, n)
	for i := 0; i < n; i++ {
		sl[i][0], sl[i][1] = ni(), ni()
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				var p1 IntPar
				var p2 IntPar
				p1[0], p1[1] = sl[i][0]-sl[j][0], sl[i][1]-sl[j][1]
				p2[0], p2[1] = sl[i][0]-sl[k][0], sl[i][1]-sl[k][1]
				t := p1[0]*p2[1] - p1[1]*p2[0]
				if t != 0 {
					count += 1
				}
			}
		}
	}
	fmt.Fprintln(wtr, count)
}

func nor(x1, y1, x2, y2 int) float64 {
	x := x1 - x2
	y := y1 - y2
	return math.Sqrt(float64(x ^ 2 + y ^ 2))
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

/*
a b c

<a b c d>

a b c
a c d

b c d

c



*/
