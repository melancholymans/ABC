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
	r1 := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(r1[0])
	k, _ := strconv.Atoi(r1[1])
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	sl := make([]int, n)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(r2[i])
		sl[i] = a
	}
	sort.Sort(sort.IntSlice(sl))
	total := 0
	for i := 0; i < k; i++ {
		total += sl[i]
	}
	fmt.Fprintln(writer, total)
}
