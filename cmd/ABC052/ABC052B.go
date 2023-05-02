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
	sc.Scan()
	s := sc.Text()[:n]
	x := 0
	maxm := math.MinInt64
	for i := 0; i < n; i++ {
		if s[i] == 73 {
			x += 1
		} else {
			x -= 1
		}
		if maxm < x {
			maxm = x
		}
	}
	if maxm <= 0 {
		maxm = 0
	}
	fmt.Fprintln(writer, maxm)
}
