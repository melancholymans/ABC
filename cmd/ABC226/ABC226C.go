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
	rst := 0
	n := ni()
	ts := make([]int, n)
	used := make([]bool, n)
	ks := make([][]int, n)
	for i := 0; i < n; i++ {
		t, k := ni2()
		ts[i] = t
		ks[i] = make([]int, k)
		for j := 0; j < k; j++ {
			ks[i][j] = ni() - 1
		}
	}
	var dfs func(i int)
	dfs = func(i int) {
		if used[i] {
			return
		}
		used[i] = true
		rst += ts[i]
		for _, v := range ks[i] {
			dfs(v)
		}
	}
	dfs(n - 1)
	fmt.Fprintln(wtr, rst)
}
