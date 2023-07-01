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
	sl := nis(n)
	mmax := math.MinInt64
	for i := 0; i < n; i++ {
		if mmax < sl[i] {
			mmax = sl[i]
		}
	}
	count := make([]int, mmax+1)
	for i := 2; i <= mmax; i++ {
		fmt.Fprintln(wtr, "i=", i)
		for j := 0; j < n; j++ {
			if sl[j]%i == 0 {
				count[i] += 1
				fmt.Fprintf(wtr, "count[%d]=%d\n", i, count[i])
			}

		}
	}
	fmt.Fprintln(wtr, count)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
