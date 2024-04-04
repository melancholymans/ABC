package main

import (
	"bufio"
	"fmt"
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
	n, k := ni2()
	ns := make([]int, n)
	nss := make([]int, n)
	for i := 0; i < n; i++ {
		a, b, c := ni3()
		ns[i] = a + b + c
		nss[i] = a + b + c
	}
	sort.Slice(nss, func(i, j int) bool {
		return nss[i] < nss[j]
	})
	nss = append(nss, math.MaxInt64)
	for i := 0; i < n; i++ {
		f := func(c int) bool {
			if nss[c] <= ns[i]+300 {
				return true
			}
			return false
		}
		r := n - bs(0, n+1, f)
		if r <= k {
			fmt.Fprintln(wtr, "Yes")
		} else {
			fmt.Fprintln(wtr, "No")
		}
	}
}

func bs(ok, ng int, f func(int) bool) int {
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		if f(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
