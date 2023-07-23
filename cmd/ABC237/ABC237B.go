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
	h, w := ni2()
	sl := make([][]int, h)
	for i := 0; i < h; i++ {
		sl[i] = make([]int, w)
		for j := 0; j < w; j++ {
			sl[i][j] = ni()
		}
	}
	r := rotate(sl)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			fmt.Fprint(wtr, r[i][j], " ")
		}
		fmt.Fprintln(wtr, " ")
	}
}

func rotate(sl [][]int) [][]int {
	h, w := len(sl), len(sl[0])
	r := make([][]int, w)
	for i := 0; i < w; i++ {
		r[i] = make([]int, h)
		for j := 0; j < h; j++ {
			r[i][j] = sl[j][i]
		}
	}
	return r
}
