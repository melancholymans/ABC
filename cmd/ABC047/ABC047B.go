package main

import (
	"bufio"
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

type IntTriple [3]int

func i2s(l, m int, def int) [][]int {
	sl := make([][]int, l)
	for i := 0; i < l; i++ {
		sl[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sl[i][j] = def
		}
	}
	return sl
}

func main() {
	defer flush()
	//o := 0
	w, h, n := ni3()
	sl := make([]IntTriple, n)
	//mp := i2s(w, h, 1)
	sm := make([][]int, w)

	for i := 0; i < n; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < h; k++ {
				switch as[i] {
				case 1:
					if j < xs[i] {
						mp[j][k] = 0
					}
				case 2:
					if j >= xs[i] {
						mp[j][k] = 0
					}
				case 3:
					if k < ys[i] {
						mp[j][k] = 0
					}
				case 4:
					if k >= ys[i] {
						mp[j][k] = 0
					}
				}

			}
		}
	}
	for j := 0; j < w; j++ {
		for k := 0; k < h; k++ {
			o += mp[j][k]
		}
	}

	out(o)
}
