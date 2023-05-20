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
	r := strings.Split(sc.Text(), " ")
	maxm := math.MinInt64
	total := 0
	for i := 0; i < n; i++ {
		l, _ := strconv.Atoi(r[i])
		if maxm < l {
			maxm = l
		}
		total += l
	}
	if (total - maxm) > maxm {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
