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
	a, b, c := ni3()
	d, e, f := ni3()
	x := ni()
	n := x / (a + c) * a
	m := x % (a + c)
	if m > a {
		m = a
	}
	taka := (n + m) * b
	n = x / (d + f) * d
	m = x % (d + f)
	if m > d {
		m = d
	}
	aoki := (n + m) * e
	if taka > aoki {
		fmt.Fprintln(wtr, "Takahashi")
	} else if taka == aoki {
		fmt.Fprintln(wtr, "Draw")
	} else {
		fmt.Fprintln(wtr, "Aoki")
	}
}
