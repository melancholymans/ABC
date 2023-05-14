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
	s := sc.Text()
	n := len(s) - 2
	minm := math.MaxInt64
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(s[i : i+3])
		tm := Abs(a - 753)
		if minm > tm {
			minm = tm
		}
	}
	fmt.Fprintln(writer, minm)
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
