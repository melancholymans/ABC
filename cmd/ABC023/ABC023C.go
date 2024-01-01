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
	r, c, k := ni3()
	n := ni()
	rcs := make([][]int, r)
	rsum := make([]int, r)
	csum := make([]int, c)
	rts := make([]int, c+1)
	cts := make([]int, r+1)
	for i := 0; i < n; i++ {
		r := ni() - 1
		c := ni() - 1
		rcs[r] = append(rcs[r], c)
		rsum[r] += 1
		csum[c] += 1
	}
	count := 0
	for i := 0; i < r; i++ {
		rts[rsum[i]]++
	}
	for i := 0; i < c; i++ {
		cts[csum[i]]++
	}
	for i := 0; i < r; i++ {
		if rsum[i] > k {
			continue
		}
		if k-rsum[i] <= r {
			count += cts[k-rsum[i]]
		}
		for _, v := range rcs[i] {
			if csum[v] == k-rsum[i] {
				count -= 1
			}
			if csum[v] == k-rsum[i]+1 {
				count += 1
			}
		}
	}
	fmt.Fprintln(wtr, count)
}
