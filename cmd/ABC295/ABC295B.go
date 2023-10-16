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
	r, c := ni2()
	sl := make([][]string, r)
	slc := make([][]string, r)
	for i := 0; i < r; i++ {
		sl[i] = strings.Split(ns(), "")
		slc[i] = make([]string, c)
		for j := 0; j < c; j++ {
			slc[i][j] = sl[i][j]
		}
	}
	for x := 0; x < r; x++ {
		for y := 0; y < c; y++ {
			d, _ := strconv.Atoi(sl[x][y])
			if d > 0 {
				for i := 0; i < r; i++ {
					for j := 0; j < c; j++ {
						if abs(x-i)+abs(y-j) <= d {
							slc[i][j] = "."
						}
					}
				}
			}
		}
	}
	for x := 0; x < r; x++ {
		for y := 0; y < c; y++ {
			fmt.Fprint(wtr, slc[x][y])
		}
		fmt.Fprintln(wtr, " ")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
