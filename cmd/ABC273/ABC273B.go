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
	rst := 0
	tx := s2i(ns())
	k := ni()
	x := make([]int, 18)
	ti := 0
	for i := len(tx) - 1; i >= 0; i-- {
		x[ti] = tx[i]
		ti += 1
	}
	for i := 0; i < k; i++ {
		if x[i] >= 5 {
			x[i+1] += 1
		}
		x[i] = 0
	}
	t := 1
	for i := 0; i < 18; i++ {
		rst += x[i] * t
		t *= 10
	}
	fmt.Fprintln(wtr, rst)
}

func s2i(s string) []int {
	sl := make([]int, len(s))
	kl := []byte(s)
	for i := 0; i < len(s); i++ {
		sl[i] = int(kl[i] - 48)
	}
	return sl
}
