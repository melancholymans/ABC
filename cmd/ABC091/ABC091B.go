package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	mp := make(map[string]int, 0)
	for i := 0; i < n; i++ {
		sc.Scan()
		s := sc.Text()
		mp[s] += 1
	}
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	for i := 0; i < m; i++ {
		sc.Scan()
		t := sc.Text()
		mp[t] -= 1
	}
	maxm := math.MinInt64
	for _, v := range mp {
		if maxm < v {
			maxm = v
		}
	}
	fmt.Fprintln(writer, Max(maxm, 0))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
