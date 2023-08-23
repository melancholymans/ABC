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
	n := ni()
	sl := make([][]string, n)
	for i := 0; i < n; i++ {
		sl[i] = strings.Split(ns(), "")
	}
	l := 1
	line := make([]string, n*4+1 /*-4*/)
	line[0] = "*"
	for i := 0; i < n; i++ {
		line[l] = sl[0][i]
		//fmt.Fprintln(wtr, " l=", l, " i=", i, " sl[][]=", sl[0][i], " line[l]=", line[l])
		l += 1
	}
	l -= 1
	for i := 0; i < n; i++ {
		line[l] = sl[i][n-1]
		//fmt.Fprintln(wtr, " l=", l, " i=", i, " sl[][]=", sl[i][n-1], " line[l]=", line[l])
		l += 1
	}
	l -= 1
	for i := n - 1; i >= 0; i-- {
		line[l] = sl[n-1][i]
		//fmt.Fprintln(wtr, " l=", l, " i=", i, " sl[][]=", sl[n-1][i], " line[l]=", line[l])
		l += 1
	}
	l -= 1
	for i := n - 1; i >= 0; i-- {
		line[l] = sl[i][0]
		//fmt.Fprintln(wtr, " l=", l, " i=", i, " sl[][]=", sl[i][0], " line[l]=", line[l])
		l += 1
	}
	fmt.Fprintln(wtr, line)
}
