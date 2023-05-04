package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	minm := math.MaxInt64
	maxm := math.MinInt64
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(r2[i])
		if a < minm {
			minm = a
		}
		if a > maxm {
			maxm = a
		}
	}
	fmt.Fprintln(writer, maxm-minm)
}
