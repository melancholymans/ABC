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
	min := math.MaxInt64
	var t int
	for i := 0; i < n; i++ {
		sc.Scan()
		t, _ = strconv.Atoi(sc.Text())
		if t < min {
			min = t
		}
	}
	fmt.Fprintln(writer, min)
}
