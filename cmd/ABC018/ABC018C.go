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

type cusum struct {
	l int
	s []int
}

func newcusum(sl []int) *cusum {
	c := &cusum{}
	c.l = len(sl)
	c.s = make([]int, len(sl)+1)
	for i, v := range sl {
		c.s[i+1] = c.s[i] + v
	}
	return c
}

func (c *cusum) getRange(f, t int) int {
	if f > t || f >= c.l {
		return 0
	}
	return c.s[t+1] - c.s[f]
}

func main() {
	defer wtr.Flush()
	total := 0
	r, c, k := ni3()
	sl := make([][]int, r)
	for i := 0; i < r; i++ {
		sl[i] = make([]int, c)
	}
	cs := make([]*cusum, r)
	for i := 0; i < r; i++ {
		s := ns()
		for j, v := range s {
			if string(v) != "o" {
				sl[i][j] = 1
			}
		}
		cs[i] = newcusum(sl[i])
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i < k-1 || r-k < i {
				continue
			}
			if j < k-1 || c-k < j {
				continue
			}
			t := 0
			for l := 0; l < k; l++ {
				ui := i - l
				di := i + l
				lj := j - (k - 1 - l)
				rj := j + (k - 1 - l)
				t += cs[ui].getRange(lj, rj)
				t += cs[di].getRange(lj, rj)

			}
			if t == 0 {
				total++
			}
		}
	}
	fmt.Fprintln(wtr, total)
}
