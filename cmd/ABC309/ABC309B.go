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
	rl := make([][]string, n)
	for i := 0; i < n; i++ {
		rl[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == n-1 || j == n-1 {
				if i == 0 && j < n-1 {
					rl[i][j+1] = sl[i][j]
				}
				if i < n-1 && j == n-1 {
					rl[i+1][j] = sl[i][j]
				}
				if i == n-1 && j > 0 {
					rl[i][j-1] = sl[i][j]
				}
				if i > 0 && j == 0 {
					rl[i-1][j] = sl[i][j]
				}
			} else {
				rl[i][j] = sl[i][j]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(wtr, rl[i][j])
		}
		fmt.Fprintln(wtr, " ")
	}
}
