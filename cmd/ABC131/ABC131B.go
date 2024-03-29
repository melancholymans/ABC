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
	n, l := ni2()
	mmin := math.MaxInt64
	b := calc(n, l)
	sig := 1
	for i := 0; i < n; i++ {
		c := b - calc2(n, l, i)
		if c < 0 {
			sig = -1
		}
		c = Abs(c)
		if c < mmin {
			mmin = c
		}
	}
	fmt.Fprintln(wtr, b-mmin*sig)
}

func calc(n, l int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += l + (i + 1) - 1
	}
	return total
}

func calc2(n, l, skip int) int {
	total := 0
	for i := 0; i < n; i++ {
		if skip == i {
			continue
		}
		total += l + (i + 1) - 1
	}
	return total
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
