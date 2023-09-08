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
	sl := make([][]int, n)
	for i := 0; i < n; i++ {
		sl[i] = make([]int, n)
		for j, s := range strings.Split(ns(), "") {
			num, _ := strconv.Atoi(s)
			sl[i][j] = num
		}
	}
	//oneNumber()
	idh, idw, rst := area(n, sl)
	for lp := 0; lp < n-1; lp++ {
		h, w := idh, idw
		c := 0
		mmax := math.MinInt64
		if w-1 < 0 {
			w = n - 1
		} else {
			w = w - 1
		}
		c = sl[h][w]
		if mmax < c {
			mmax = c
			idw = w - 1
		}
		if w+1 > n-1 {
			w = 0
		}
		c = sl[h][w+1]
		if mmax < c {
			mmax = c
			idw = w + 1
		}
		c = sl[h-1][w]
		if mmax < c {
			mmax = c
			idh = h - 1
		}
		c = sl[h+1][w]
		if mmax < c {
			mmax = c
			idh = h + 1
		}
		rst += strconv.Itoa(mmax)
	}
	fmt.Fprintln(wtr, rst)
}

func area(n int, sl [][]int) (int, int, string) {
	mmax := math.MinInt64
	idh, idw := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mmax < sl[i][j] {
				mmax = sl[i][j]
				idh = i
				idw = j
			}
		}
	}
	return idh, idw, strconv.Itoa(mmax)
}

/*
func oneNumber(){

}
*/

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
