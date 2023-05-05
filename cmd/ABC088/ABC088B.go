package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	total := 0
	l := len(a)
	ng := 1
	for i := 0; i < l; i++ {
		total += a[i] * ng
		ng = -ng
	}
	fmt.Println(total)
}
