package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func main() {
	defer wtr.Flush()
	n, k, x := ni3()
	ns := nis(n)
	dc := 0
	total := 0
	for i, v := range ns {
		total += v
		dc += v / x
		ns[i] = v % x
	}
	if dc >= k {
		total -= k * x
	} else {
		total -= dc * x
		k -= dc
		sort.Slice(ns, func(i, j int) bool {
			return ns[i] > ns[j]
		})
		for i := 0; i < min(k, n); i++ {
			total -= ns[i]
		}
	}
	fmt.Fprintln(wtr, total)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
