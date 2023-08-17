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

type place struct {
	d string
	p int
}

func main() {
	defer wtr.Flush()
	n, m := ni2()
	c := make([]string, n)
	for i := 0; i < n; i++ {
		c[i] = ns()
	}
	mp := map[string]int{}
	d := make([]string, m)
	for i := 0; i < m; i++ {
		d[i] = ns()
	}
	p0 := ni()
	p := make([]int, m)
	for i := 0; i < m; i++ {
		p[i] = ni()
	}
	for i := 0; i < m; i++ {
		mp[d[i]] = p[i]
	}
	count := 0
	for i := 0; i < n; i++ {
		if v, ok := mp[c[i]]; ok {
			count += v
		} else {
			count += p0
		}
	}
	fmt.Fprintln(wtr, count)
}
