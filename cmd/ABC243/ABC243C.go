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

type point struct {
	x int
	y int
	m string
}

func main() {
	defer wtr.Flush()
	n := ni()
	xs := make([]int, n)
	ys := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = ni()
		ys[i] = ni()
	}
	s := ns()
	yrm := make(map[int]int)
	ylm := make(map[int]int)
	for i, v := range s {
		if string(v) == "R" {
			if v, ok := yrm[ys[i]]; ok {
				yrm[ys[i]] = min(v, xs[i])
			} else {
				yrm[ys[i]] = xs[i]
			}
		} else {
			if v, ok := ylm[ys[i]]; ok {
				ylm[ys[i]] = max(v, xs[i])
			} else {
				ylm[ys[i]] = xs[i]
			}
		}
		if v1, ok := yrm[ys[i]]; ok {
			if v2, ok := ylm[ys[i]]; ok {
				if v2 > v1 {
					fmt.Fprintln(wtr, "Yes")
					return
				}
			}
		}
	}
	fmt.Fprintln(wtr, "No")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
