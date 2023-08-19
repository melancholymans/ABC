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
	n, d := ni2()
	sl := make([][]int, n)
	for i := 0; i < n; i++ {
		sl[i] = make([]int, d)
		s := strings.Split(ns(), "")
		for j := 0; j < d; j++ {
			if s[j] == "o" {
				sl[i][j] = 1
			}
		}
	}
	cum := make([]int, d)
	for i := 0; i < d; i++ {
		for j := 0; j < n; j++ {
			cum[i] += sl[j][i]
		}
	}
	mmax := 0
	count := 0
	for i := 0; i < d; i++ {
		if cum[i] == n {
			count += 1
		}
		if cum[i] != n {
			count = 0
		}
		if mmax < count {
			mmax = count
		}
	}
	fmt.Fprintln(wtr, mmax)
}
