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
	h, m := ni2()
	for {
		h1 := h / 10
		h2 := h % 10
		m1 := m / 10
		m2 := m % 10
		if isMistake(h1, h2, m1) {
			fmt.Fprintf(wtr, "%02d %02d", h1*10+h2, m1*10+m2)
			return
		}
		h, m = timeCalc(h1*10+h2, m1*10+m2)
	}
}

func isMistake(h1, h2, m1 int) bool {
	//見間違えやすい表示ならtrueを返す
	if h1 == 2 {
		if m1 >= 0 && m1 <= 3 && (h2 >= 0 && h2 <= 5) {
			return true
		}
	} else {
		if (m1 >= 0 && m1 <= 9) && (h2 >= 0 && h2 <= 5) {
			return true
		}
	}
	return false
}

func timeCalc(h, m int) (int, int) {
	m = m + 1
	if m/60 == 1 {
		m = 0
		h += 1
	}
	if h/24 == 1 {
		h = 0
	}
	return h, m
}
