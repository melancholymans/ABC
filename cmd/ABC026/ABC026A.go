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
	r := sc.Text()
	a, _ := strconv.Atoi(r)
	max := math.MinInt64
	for x := 0; x <= a; x++ {
		z := x * (a - x)
		if z > max {
			max = z
		}
	}
	fmt.Fprintln(writer, max)
}
