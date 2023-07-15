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
	h, w, x, y := ni4()
	x, y = x-1, y-1
	sl := make([][]string, h)
	for i := 0; i < h; i++ {
		sl[i] = make([]string, w)
		r := strings.Split(ns(), "")
		for j := 0; j < w; j++ {
			sl[i][j] = r[j]
		}
	}
	count := 0
	//自座標
	if sl[x][y] == "." {
		count += 1
	}
	//横方向
	for i := y + 1; i < w; i++ {
		if sl[x][i] == "#" {
			break
		} else if sl[x][i] == "." {
			count += 1
		}
	}
	for i := y - 1; i >= 0; i-- {
		if sl[x][i] == "#" {
			break
		} else if sl[x][i] == "." {
			count += 1
		}
	}
	//縦方向
	for i := x + 1; i < h; i++ {
		if sl[i][y] == "#" {
			break
		} else if sl[i][y] == "." {
			count += 1
		}
	}
	for i := x - 1; i >= 0; i-- {
		if sl[i][y] == "#" {
			break
		} else if sl[i][y] == "." {
			count += 1
		}
	}
	fmt.Fprintln(wtr, count)
}
