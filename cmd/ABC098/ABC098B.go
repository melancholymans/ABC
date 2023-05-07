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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := sc.Text()
	ma := make(map[string]int)
	mb := make(map[string]int)
	maxm := math.MinInt64
	for i := 1; i < n; i++ {
		sa := strings.Split(s[:i], "")
		sb := strings.Split(s[i:], "")
		for _, v := range sa {
			ma[v] += 1
		}
		for _, v := range sb {
			mb[v] += 1
		}
		count := 0
		for k, _ := range ma {
			if v, ok := mb[k]; ok && v > 0 {
				count += 1
			}
		}
		if count > maxm {
			maxm = count
		}
		for k, _ := range ma {
			ma[k] = 0
		}
		for k, _ := range mb {
			mb[k] = 0
		}
	}
	fmt.Fprintln(writer, maxm)
}
