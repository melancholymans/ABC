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
	rst := 0
	for i := 0; i < n; i++ {
		sl[i] = make([]int, n)
		for j, s := range strings.Split(ns(), "") {
			num, _ := strconv.Atoi(s)
			sl[i][j] = num
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := []int{1, 1, 0, -1, -1, -1, 0, 1}
			y := []int{0, 1, 1, 1, 0, -1, -1, -1}
			for k := 0; k < 8; k++ {
				t := sl[i][j]
				for l := 1; l <= n-1; l++ {
					t *= 10
					t += sl[(i+x[k]*l+n)%n][(j+y[k]*l+n)%n]
				}
				rst = Max(rst, t)
			}
		}
	}
	fmt.Fprintln(wtr, rst)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
