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
	for i1 := 0; i1 < h; i1++ {
		for i2 := i1 + 1; i2 < h; i2++ {
			for j1 := 0; j1 < w; j1++ {
				for j2 := j1 + 1; j2 < w; j2++ {
					if sl[i1][j1]+sl[i2][j2] > sl[i2][j1]+sl[i1][j2] {
						fmt.Fprintln(wtr, "No")
						return
					}
				}
			}
		}
	}
	fmt.Fprintln(wtr, "Yes")
}
