package main

import (
	"bufio"
	"fmt"
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
	r1 := sc.Text()
	n, _ := strconv.Atoi(r1)
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	sl := make([]int, 0)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(r2[i])
		sl = append(sl, a)
	}
	count := 0
	for i := 0; i < n-3+1; i++ {
		slc := make([]int, 3)
		copy(slc, sl[i:i+3])
		sort.Sort(sort.IntSlice(slc))
		if slc[1] == sl[i+1] {
			count += 1
		}
	}
	fmt.Fprintln(writer, count)
}
