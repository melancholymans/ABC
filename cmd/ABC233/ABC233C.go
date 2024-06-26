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
	n, x := ni2()
	ns := make([][]int, n)
	for i := 0; i < n; i++ {
		l := ni()
		ns[i] = make([]int, l)
		for j := 0; j < l; j++ {
			ns[i][j] = ni()
		}
	}
	count := 0
	var dfs func(i, v int)
	dfs = func(i, v int) {
		for j := 0; j < len(ns[i]); j++ {
			if v%ns[i][j] != 0 {
				continue
			}
			if i == n-1 {
				if v == ns[i][j] {
					count += 1
				}
			} else {
				dfs(i+1, v/ns[i][j])
			}
		}
	}
	dfs(0, x)
	fmt.Fprintln(wtr, count)
}
