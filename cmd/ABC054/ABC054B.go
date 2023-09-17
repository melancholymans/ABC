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

func stois(s string, baseRune rune) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[i] = int(v - baseRune)
	}
	return r
}

func main() {
	defer wtr.Flush()
	n, m := ni2()
	nm := make([][]int, n)
	mm := make([][]int, m)
	for i := 0; i < n; i++ {
		nm[i] = stois(ns(), 'a')
	}
	for i := 0; i < m; i++ {
		mm[i] = stois(ns(), 'a')
	}
	for i := 0; i < n-m+1; i++ {
		for j := 0; j < n-m+1; j++ {
			isSame := true
			for k := 0; k < m; k++ {
				for l := 0; l < m; l++ {
					if mm[k][l] != nm[i+k][j+l] {
						isSame = false
					}
				}
			}
			if isSame {
				fmt.Fprintln(wtr, "Yes")
				return
			}
		}
	}
	fmt.Fprintln(wtr, "No")
}
