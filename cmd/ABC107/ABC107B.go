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
	h, w := ni2()
	sl := make([][]string, h)
	for i := 0; i < h; i++ {
		sl[i] = make([]string, w)
		s := strings.Split(ns(), "")
		for j := 0; j < w; j++ {
			sl[i][j] = s[j]
		}
	}
	cp1 := make([][]string, 0)
	for i := 0; i < h; i++ {
		if calc(sl[i]) {
			cp1 = append(cp1, sl[i])
		}
	}
	r := rotate(cp1)
	cp2 := make([][]string, 0)
	for i := 0; i < len(r); i++ {
		if calc(r[i]) {
			cp2 = append(cp2, r[i])
		}
	}
	r = rotate(cp2)
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[0]); j++ {
			fmt.Fprint(wtr, r[i][j])
		}
		fmt.Fprintln(wtr, " ")
	}
}

func calc(s []string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == "#" {
			return true
		}
	}
	return false
}

func rotate(sl [][]string) [][]string {
	h, w := len(sl), len(sl[0])
	r := make([][]string, w)
	for i := 0; i < w; i++ {
		r[i] = make([]string, h)
		for j := 0; j < h; j++ {
			r[i][j] = sl[j][i]
		}
	}
	return r
}
