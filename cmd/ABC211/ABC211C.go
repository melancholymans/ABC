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
	s := ns()
	dp := make([]int, 8)
	for _, v := range s {
		switch string(v) {
		case "c":
			dp[0]++
		case "h":
			dp[1] = madd(dp[1], dp[0])
		case "o":
			dp[2] = madd(dp[2], dp[1])
		case "k":
			dp[3] = madd(dp[3], dp[2])
		case "u":
			dp[4] = madd(dp[4], dp[3])
		case "d":
			dp[5] = madd(dp[5], dp[4])
		case "a":
			dp[6] = madd(dp[6], dp[5])
		case "i":
			dp[7] = madd(dp[7], dp[6])
		}
	}
	fmt.Fprintln(wtr, dp[7])
}

func madd(a, b int) int {
	const mod = 1000000007
	a += b
	if a < 0 {
		a += mod
	} else if a >= mod {
		a -= mod
	}
	return a
}
