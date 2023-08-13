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
	l := len(sl)
	total := sum(sl)
	if total == 0 {
		fmt.Fprintln(wtr, 0)
		return
	}
	if total%l != 0 {
		fmt.Fprintln(wtr, -1)
		return
	}
	k := total / l
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = sl[i] - k
	}
	cum := make([]int, n+1)
	cum[0] = 0
	for i := 1; i <= n; i++ {
		cum[i] = cum[i-1] + p[i-1]
	}
	cum = cum[1:]
	gup, bge := 0, 0
	for i := 0; i < n; i++ {
		gup += 1
		if cum[i] == 0 {
			bge += gup - 1
			gup = 0
		}
	}
	fmt.Fprintln(wtr, bge)
}

func sum(sl []int) int {
	total := 0
	for i := 0; i < len(sl); i++ {
		total += sl[i]
	}
	return total
}
