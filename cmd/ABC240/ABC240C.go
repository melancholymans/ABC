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
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, 10001)
	}
	dp[0][0] = true
	for i := 0; i < n; i++ {
		a, b := ni2()
		for j := 0; j < 10000; j++ {
			if !dp[i][j] {
				continue
			}
			dp[i+1][min(10000, j+a)] = true
			dp[i+1][min(10000, j+b)] = true
		}
	}
	if dp[n][x] {
		fmt.Fprintln(wtr, "Yes")
		return
	} else {
		fmt.Fprintln(wtr, "No")
		return
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
