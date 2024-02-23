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

func ns() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	defer wtr.Flush()
	n := ni()
	mf := make([][]bool, n)
	mt := make([][]bool, n)
	mf4d := make([][][]bool, 4)
	var xmin, xmax, ymin, ymax int
	xmin, ymin = n, n
	fc := 0
	for i := 0; i < n; i++ {
		s := ns()
		mf[i] = make([]bool, n)
		for j, v := range s {
			if string(v) == "#" {
				mf[i][j] = true
				xmin = min(xmin, i)
				xmax = max(xmax, i)
				ymin = min(ymin, j)
				ymax = max(ymax, j)
				fc += 1
			}
		}
	}
	ft := 0
	for i := 0; i < n; i++ {
		s := ns()
		mt[i] = make([]bool, n)
		for j, v := range s {
			if string(v) == "#" {
				mt[i][j] = true
				ft += 1
			}
		}
	}
	if fc != ft {
		fmt.Fprintln(wtr, "No")
		return
	}
	rx := xmax - xmin + 1
	ry := ymax - ymin + 1
	mf4d[0] = make([][]bool, rx)
	mf4d[1] = make([][]bool, rx)
	mf4d[2] = make([][]bool, ry)
	mf4d[3] = make([][]bool, ry)
	for i := 0; i < rx; i++ {
		mf4d[0][i] = make([]bool, ry)
		mf4d[1][i] = make([]bool, ry)
	}
	for j := 0; j < ry; j++ {
		mf4d[2][j] = make([]bool, rx)
		mf4d[3][j] = make([]bool, rx)
	}
	for i := 0; i < rx; i++ {
		for j := 0; j < ry; j++ {
			mf4d[0][i][j] = mf[i+xmin][j+ymin]
			mf4d[1][rx-i-1][ry-j-1] = mf[i+xmin][j+ymin]
			mf4d[2][ry-j-1][i] = mf[i+xmin][j+ymin]
			mf4d[3][j][rx-i-1] = mf[i+xmin][j+ymin]
		}
	}
	for _, mf4di := range mf4d {
		for i := 0; i <= n-len(mf4di); i++ {
			for j := 0; j <= n-len(mf4di[0]); j++ {
				r := 0
				for k, v := range mf4di {
					for l, v2 := range v {
						if mt[i+k][j+l] == v2 {
							r += 1
						}
					}
				}
				if r == rx*ry {
					fmt.Fprintln(wtr, "Yes")
					return
				}
			}
		}
	}
	fmt.Fprintln(wtr, "No")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
