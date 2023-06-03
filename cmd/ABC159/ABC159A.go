package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const mod1000000007 = 1000000007
const mod = mod1000000007

type combFactorial struct {
	fac    []int
	facinv []int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(r[0])
	m, _ := strconv.Atoi(r[1])
	c := NewcombFactorial(Max(n, m) + 1)
	fmt.Println(c.Combination(n, 2) + c.Combination(m, 2))
}

func NewcombFactorial(n int) *combFactorial {
	fac := make([]int, n)
	facinv := make([]int, n)
	fac[0] = 1
	facinv[0] = minvfermat(1, mod)
	for i := 1; i < n; i++ {
		fac[i] = mmul(i, fac[i-1])
		facinv[i] = minvfermat(fac[i], mod)
	}
	return &combFactorial{
		fac:    fac,
		facinv: facinv,
	}
}

func (c *combFactorial) Combination(n, r int) int {
	if r > n {
		return 0
	}
	return mmul(mmul(c.fac[n], c.facinv[r]), c.facinv[n-r])
}

func minvfermat(a, m int) int {
	return mpow(a, m-2, mod)
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

func mmul(a, b int) int {
	return a * b % mod
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
