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

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r1 := sc.Text()
	n, _ := strconv.Atoi(r1)
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	sl := make([]int, 0)
	max := math.MinInt64
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(r2[i])
		sl = append(sl, a)
		if a > max {
			max = a
		}
	}
	sort.Ints(sl)
	l := len(sl)
	count := 0
	for k := 1; k <= max; k++ {
		idx := LowerBound(k, sl)
		if idx+1 == l-idx-1 {
			count += 1
		}
	}
	fmt.Fprintln(writer, count)
}

func LowerBound(v int, sl []int) int {
	if len(sl) == 0 {
		return -1
	}
	idx := bs(0, len(sl)-1, func(c int) bool {
		return sl[c] < v
	})
	return idx
}

func bs(low, hi int, f func(int) bool) int {
	if !f(low) {
		return -1
	}
	if f(hi) {
		return hi
	}
	for Abs(low-hi) > 1 {
		mid := (low + hi) / 2
		if f(mid) {
			low = mid
		} else {
			hi = mid
		}
	}
	return low
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
