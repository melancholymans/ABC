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
	rst := 0
	n := ni()
	rst = mdiv(mmul((n+1)%mod998244353, n%mod998244353), 2)
	t := 10
	for {
		if n < t {
			break
		}
		rst = madd(rst, -mmul((t-1)%mod998244353, min(n-t+1, t*9)%mod998244353))
		t *= 10
	}
	fmt.Fprintln(wtr, rst)
}

func mdiv(a, b int) int {
	a %= mod998244353
	return a * minvfermat(b, mod998244353) % mod998244353
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

func mmul(a, b int) int {
	return a * b % mod998244353
}

func mpow(a, n, m int) int {
	if m == 1 {
		return 0
	}
	r := 1
	for n > 0 {
		if n&1 == 1 {
			r = r * a % m
		}
		a, n = a*a%m, n>>1
	}
	return r
}

func minvfermat(a, m int) int {
	return mpow(a, m-2, mod998244353)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
