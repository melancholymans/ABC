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
	sum := 0
	bs := [2][3]int{}
	cs := [3][2]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			bs[i][j] = ni()
			sum += bs[i][j]
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			cs[i][j] = ni()
			sum += cs[i][j]
		}
	}
	mp := [3][3]int{}
	var dfs func(mp [3][3]int, c int) int
	dfs = func(mp [3][3]int, c int) int {
		if c == 9 {
			t := 0
			for i := 0; i < 2; i++ {
				for j := 0; j < 3; j++ {
					if mp[i][j] == mp[i+1][j] {
						t += bs[i][j]
					}
				}
			}
			for i := 0; i < 3; i++ {
				for j := 0; j < 2; j++ {
					if mp[i][j] == mp[i][j+1] {
						t += cs[i][j]
					}
				}
			}
			return t
		}
		t := 0
		if c%2 == 1 {
			t = math.MaxInt64
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if mp[i][j] != 0 {
					continue
				}
				mp[i][j] = c%2 + 1
				r := dfs(mp, c+1)
				if c%2 == 1 {
					t = min(t, r)
				} else {
					t = max(t, r)
				}
				mp[i][j] = 0
			}
		}
		return t
	}
	total := dfs(mp, 0)
	fmt.Fprintln(wtr, total)
	fmt.Fprintln(wtr, sum-total)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
