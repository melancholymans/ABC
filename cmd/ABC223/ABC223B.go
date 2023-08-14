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
	s := []byte(ns())
	mmax := shift_left(s, 0)
	mmin := shift_left(s, 0)
	for i := 1; i < len(s); i++ {
		ht := shift_left(s, i)
		if mmax < ht {
			mmax = ht
		}
		if mmin > ht {
			mmin = ht
		}
	}
	fmt.Fprintln(wtr, mmin)
	fmt.Fprintln(wtr, mmax)
}

func shift_left(s []byte, n int) string {
	l := len(s)
	r := make([]byte, l)
	for i := 0; i < l-n; i++ {
		r[i] = s[i+n]
	}
	for i := 0; i < n; i++ {
		r[i+l-n] = s[i]
	}
	return string(r)
}
