package main

import (
	"bufio"
	"fmt"
	"os"
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
	m, _ := strconv.Atoi(r1[1])
	x, _ := strconv.Atoi(r1[2])

	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	sm := make([]int, m)
	for i := 0; i < m; i++ {
		a, _ := strconv.Atoi(r2[i])
		sm[i] = a
	}
	sl := make([]int, n+1)
	for _, v := range sm {
		sl[v] = 1
	}
	//left total
	lt := 0
	for i := 0; i < x; i++ {
		lt += sl[i]
	}
	//right total
	rt := 0
	for i := x; i < n; i++ {
		rt += sl[i]
	}
	if lt > rt {
		fmt.Fprintln(writer, rt)
	} else {
		fmt.Fprintln(writer, lt)
	}
}
