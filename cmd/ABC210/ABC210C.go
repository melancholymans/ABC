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
	count := 0
	n, k := ni2()
	sl := nis(n)
	mp := make(map[int]int)
	t := 0
	for i := 0; i < n; i++ {
		c := sl[i]
		if mp[c] == 0 {
			t++
		}
		mp[c] += 1
		if i >= k {
			cb := sl[i-k]
			mp[cb] -= 1
			if mp[cb] == 0 {
				t--
			}
		}
		count = mmax(count, t)
	}
	fmt.Fprintln(wtr, count)
}

func mmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
