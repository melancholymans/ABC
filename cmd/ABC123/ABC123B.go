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
	sl := [5]int{}
	sc.Scan()
	sl[0], _ = strconv.Atoi(sc.Text())
	sc.Scan()
	sl[1], _ = strconv.Atoi(sc.Text())
	sc.Scan()
	sl[2], _ = strconv.Atoi(sc.Text())
	sc.Scan()
	sl[3], _ = strconv.Atoi(sc.Text())
	sc.Scan()
	sl[4], _ = strconv.Atoi(sc.Text())
	maxm := math.MinInt64
	total := 0
	n := 0
	md := 0
	for _, v := range sl {
		if v%10 == 0 {
			md = 0
			n = v
		} else {
			md = 10 - v%10
			n = v + (10 - v%10)
		}
		if maxm < md {
			maxm = md
		}
		total += n
	}
	fmt.Fprintln(writer, total-maxm)
}
