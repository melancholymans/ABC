package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(r[0])
	k, _ := strconv.Atoi(r[1])
	sl := make([]int, 0)
	for i := 0; i < n; i++ {
		sc.Scan()
		h, _ := strconv.Atoi(sc.Text())
		sl = append(sl, h)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sl)))
	min := math.MaxInt64
	for i := 0; i < len(sl)-k+1; i++ {
		m := sl[i] - sl[i+k-1]
		if m < min {
			min = m
		}
	}
	fmt.Fprintln(writer, min)
}
