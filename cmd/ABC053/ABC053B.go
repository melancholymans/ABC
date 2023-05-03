package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := sc.Text()
	var count int
	maxm := math.MinInt64
	for i := 0; i < len(s); {
		if s[i] == 65 {
			count = 0
			for j := i; j < len(s); j++ {
				if s[j] == 90 && (j+1 == len(s) || s[j+1] != 90) {
					break
				}
				count += 1
			}
			i += count
		}
		if count > maxm {
			maxm = count
		}
		i += 1
	}
	fmt.Fprintln(writer, maxm+1)
}
