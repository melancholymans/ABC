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
	maxm := math.MinInt64
	total := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		total += p
		if maxm < p {
			maxm = p
		}
	}
	fmt.Fprintln(writer, total-maxm/2)
}
