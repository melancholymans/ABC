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

func pow(a int) int {
	return a * a
}

func main() {
	defer wtr.Flush()
	n, k := ni2()
	sa := nis(k)
	sl := make([]IntPair, n)
	for i := 0; i < n; i++ {
		sl[i][0] = ni()
		sl[i][1] = ni()
	}
	rst := 0
	for i := 0; i < n; i++ {
		mmin := math.MaxInt64
		for j := 0; j < k; j++ {
			c := pow(sl[i][0]-sl[sa[j]-1][0]) + pow(sl[i][1]-sl[sa[j]-1][1])
			if mmin > c {
				mmin = c
			}
		}
		if rst < mmin {
			rst = mmin
		}
	}
	fmt.Fprintln(wtr, math.Sqrt(float64(rst)))
}
