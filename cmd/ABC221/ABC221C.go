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
	rst := 0
	s := ns()
	n := len(s)
	sl := make([]int, n)
	for i := 0; i < n; i++ {
		sl[i] = atoi(string(s[i]))
	}
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
	for {
		for j := 1; j < n; j++ {
			a := 0
			b := 0
			if sl[0] == 0 || sl[j] == 0 {
			} else {
				for i := 0; i < n; i++ {
					if i >= j {
						a = a * 10
						a = a + sl[i]
					} else {
						b = b * 10
						b = b + sl[i]
					}
				}
			}
			rst = max(rst, a*b)
		}
		if !nextPermutation(sort.IntSlice(sl)) {
			break
		}
	}
	fmt.Fprintln(wtr, rst)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}
