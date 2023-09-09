package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
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

type IntPair [5]int

func main() {
	defer wtr.Flush()
	n, x, y, z := ni4()
	sa := nis(n)
	sb := nis(n)
	sl := make([]IntPair, n)
	for i := 0; i < n; i++ {
		sl[i][0] = i + 1
		sl[i][1] = sa[i]
		sl[i][2] = sb[i]
		sl[i][3] = sa[i] + sb[i]
		sl[i][4] = 0
	}
	//数学
	sort.Slice(sl, func(i, j int) bool {
		if sl[i][1] == sl[j][1] {
			return sl[i][0] < sl[j][0]
		}
		return sl[i][1] > sl[j][1]
	})
	for i := 0; i < x; i++ {
		sl[i][4] = 1
	}
	//英語
	sort.Slice(sl, func(i, j int) bool {
		if sl[i][4] == sl[j][4] {
			if sl[i][2] == sl[j][2] {
				return sl[i][0] < sl[j][0]
			}
			return sl[i][2] > sl[j][2]
		}
		return sl[i][4] < sl[j][4]
	})
	for i := 0; i < y; i++ {
		sl[i][4] = 1
	}
	//合計
	sort.Slice(sl, func(i, j int) bool {
		if sl[i][4] == sl[j][4] {
			if sl[i][3] == sl[j][3] {
				return sl[i][0] < sl[j][0]
			}
			return sl[i][3] > sl[j][3]
		}
		return sl[i][4] < sl[j][4]
	})
	for i := 0; i < z; i++ {
		sl[i][4] = 1
	}
	sort.Slice(sl, func(i, j int) bool {
		return sl[i][0] < sl[j][0]
	})
	for i := 0; i < len(sl); i++ {
		if sl[i][4] == 1 {
			fmt.Fprintln(wtr, sl[i][0])
		}
	}
}
