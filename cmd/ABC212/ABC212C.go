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
	n, m := ni2()
	sa := nis(n)
	sb := nis(m)
	sort.Slice(sa, func(i, j int) bool {
		return sa[i] < sa[j]
	})
	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})
	mmin := math.MaxInt64
	if sa[n-1] < sb[0] {
		mmin = abs(sa[n-1] - sb[0])
		fmt.Fprintln(wtr, mmin)
		return
	}
	if sa[0] > sb[m-1] {
		mmin = abs(sa[0] - sb[m-1])
		fmt.Fprintln(wtr, mmin)
		return
	}
	var j int
	j = 0
	for i := 0; i < n; i++ {
		for {
			mmin = min(mmin, abs(sa[i]-sb[j]))
			if sa[i] < sb[j] || j == m-1 {
				break
			}
			j++
		}
	}
	fmt.Fprintln(wtr, mmin)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
