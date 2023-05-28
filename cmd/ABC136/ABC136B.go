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
	n, _ := strconv.Atoi(r)
	total := 0
	for i := 1; i <= len(r); i++ {
		if i%2 == 0 {
			continue
		}
		if i == len(r) {
			total += n - Pow(10, len(r)-1) + 1
		} else {
			total += 9 * Pow(10, i-1)
		}
	}
	fmt.Fprintln(writer, total)
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
