package main

import (
	"bufio"
	"fmt"
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
		b, e := os.ReadFile("./input")
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

func nf() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func ns() string {
	sc.Scan()
	return sc.Text()
}

type line struct {
	a float64
	b float64
}

func main() {
	defer wtr.Flush()
	n := ni()
	sl := make([]line, n)
	totaltime := 0.0
	rst := 0.0
	for i := 0; i < n; i++ {
		sl[i].a = nf()
		sl[i].b = nf()
		totaltime += sl[i].a / sl[i].b
	}
	totaltime = totaltime / 2
	for i := 0; i < n; i++ {
		t := sl[i].a / sl[i].b
		if totaltime > t {
			totaltime -= t
			rst += sl[i].a
		} else {
			rst += sl[i].b * totaltime
			break
		}
	}
	fmt.Fprintln(wtr, rst)
}
