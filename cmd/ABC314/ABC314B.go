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

type IntPair [2]int

func main() {
	defer wtr.Flush()
	n := ni()
	sl := make([][]int, n)
	for i := 0; i < n; i++ {
		sl[i] = make([]int, 37+1)
	}
	for i := 0; i < n; i++ {
		c := ni()
		for j := 0; j < c; j++ {
			sl[i][ni()] = 1
		}
		sl[i][37] = c
	}
	x := ni()
	k := 0
	mmin := math.MaxInt64
	for i := 0; i < n; i++ {
		if sl[i][x] > 0 {
			if mmin >= sl[i][37] {
				mmin = sl[i][37]
			}
		}
	}
	rst := ""
	for i := 0; i < n; i++ {
		if sl[i][x] > 0 && sl[i][37] == mmin {
			k += 1
			rst = rst + strconv.Itoa(i+1) + " "
		}
	}
	fmt.Fprintln(wtr, k)
	fmt.Fprintln(wtr, rst)
}
