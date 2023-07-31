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
	n, m, t := ni3()
	n = n * 2
	mmax := n
	tn := 0
	for i := 0; i < m; i++ {
		a, b := ni2()
		n -= (a - tn) * 2
		if n <= 0 {
			fmt.Fprintln(wtr, "No")
			return
		}
		n += (b - a) * 2
		n = Min(n, mmax)
		tn = b
	}
	n -= (t - tn) * 2
	if n <= 0 {
		fmt.Fprintln(wtr, "No")
		return
	}
	fmt.Fprintln(wtr, "Yes")
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
