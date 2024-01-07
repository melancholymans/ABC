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
	rst := ""
	n, q := ni2()
	ls, rs := ni2s(q, 1)
	dp := make([]int, n+1)
	for i := 0; i < q; i++ {
		dp[ls[i]]++
		dp[rs[i]+1]--
	}
	t := 0
	for i := 0; i < n; i++ {
		t += dp[i]
		if t%2 == 0 {
			rst += "0"
		} else {
			rst += "1"
		}
	}
	fmt.Fprintln(wtr, rst)
}

func ni2s(n, t int) ([]int, []int) {
	sa := make([]int, n)
	sb := make([]int, n)
	for i := 0; i < n; i++ {
		sa[i], sb[i] = ni2()
		sa[i] -= t
		sb[i] -= t

	}
	return sa, sb
}
