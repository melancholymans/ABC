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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	r1 := strings.Split(sc.Text(), " ")
	slt := make([]int, n)
	total := 0
	for i := 0; i < n; i++ {
		t, _ := strconv.Atoi(r1[i])
		slt[i] = t
		total += t
	}
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	for i := 0; i < m; i++ {
		sc.Scan()
		r2 := strings.Split(sc.Text(), " ")
		p, _ := strconv.Atoi(r2[0])
		x, _ := strconv.Atoi(r2[1])
		fmt.Fprintln(writer, total-(slt[p-1]-x))
	}
}
