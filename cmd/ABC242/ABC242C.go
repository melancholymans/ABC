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

const mod998244353 = 998244353

func main() {
	defer wtr.Flush()
	n := ni()
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 10)
	}
	for i := 0; i <= 9; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 1; j < 10; j++ {
			if j != 1 {
				dp[i+1][j-1] = madd(dp[i+1][j-1], dp[i][j])
			}
			if j != 9 {
				dp[i+1][j+1] = madd(dp[i+1][j+1], dp[i][j])
			}
			dp[i+1][j] = madd(dp[i+1][j], dp[i][j])
		}
	}
	rst := 0
	for j := 1; j < 10; j++ {
		rst = madd(rst, dp[n-1][j])
	}
	fmt.Fprintln(wtr, rst)
}

func madd(a, b int) int {
	a += b
	if a < 0 {
		a += mod998244353
	} else if a >= mod998244353 {
		a -= mod998244353
	}
	return a
}
