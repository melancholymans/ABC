package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
		b, e := ioutil.ReadFile("./input")
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
func ni4() (int, int, int, int) {
	return ni(), ni(), ni(), ni()
}
func ns() string {
	sc.Scan()
	return sc.Text()
}
func nf() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func main() {
	defer wtr.Flush()
	n, a, b := ni3()
	p, q, r, s := ni4()
	h := q - p + 1
	w := s - r + 1
	k1f := max(1-a, 1-b)
	k1t := min(n-a, n-b)
	k2f := max(1-a, b-n)
	k2t := min(n-a, b-1)
	for i := 0; i < h; i++ {
		rs := make([]string, w)
		for j := 0; j < w; j++ {
			x := p + i
			y := r + j
			if x-a == y-b && x-a >= k1f && x-a <= k1t {
				rs[j] = "#"
			} else if x-a == b-y && x-a >= k2f && x-a <= k2t {
				rs[j] = "#"
			} else {
				rs[j] = "."
			}
		}
		fmt.Fprintln(wtr, strings.Join(rs, ""))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
