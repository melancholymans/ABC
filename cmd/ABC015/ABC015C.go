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
	n, k := ni2()
	sl := make([][]int, n)
	for i := 0; i < n; i++ {
		sl[i] = make([]int, k)
		for j := 0; j < k; j++ {
			sl[i][j] = ni()
		}
	}
	flag := true
	var dfs func(i, t int)
	dfs = func(i, t int) {
		if i == n {
			if t == 0 {
				flag = false
			}
			return
		}
		for j := 0; j < k; j++ {
			dfs(i+1, t^sl[i][j])
		}
	}
	dfs(0, 0)
	if flag {
		fmt.Fprintln(wtr, "Nothing")
	} else {
		fmt.Fprintln(wtr, "Found")
	}
}
