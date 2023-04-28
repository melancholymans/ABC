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
	max := math.MinInt64
	total := 0
	var name string
	for i := 0; i < n; i++ {
		sc.Scan()
		r2 := strings.Split(sc.Text(), " ")
		s := r2[0]
		p, _ := strconv.Atoi(r2[1])
		total += p
		if max < p {
			name = s
			max = p
		}
	}
	if max*2 > total {
		fmt.Fprintln(writer, name)
	} else {
		fmt.Fprintln(writer, "atcoder")
	}
}
