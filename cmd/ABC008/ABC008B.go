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
	m := map[string]int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		m[sc.Text()] += 1
	}
	max := math.MinInt64
	mname := ""
	for name, v := range m {
		if max < v {
			mname = name
			max = v
		}
	}
	fmt.Fprintln(writer, mname)
}
