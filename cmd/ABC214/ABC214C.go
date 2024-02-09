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
	n := ni()
	sl := make([]IntPar, n)
	for i := 0; i < n; i++ {
		sl[i][0] = ni()
	}
	for i := 0; i < n; i++ {
		sl[i][1] = ni()
	}
	for i := 0; i < 2*n; i++ {
		from := i % n
		next := (i + 1) % n
		sl[next][1] = min(sl[next][1], sl[from][1]+sl[from][0])
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(wtr, sl[i][1])
	}

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
