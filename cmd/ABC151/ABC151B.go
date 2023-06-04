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
	k, _ := strconv.Atoi(r1[1])
	m, _ := strconv.Atoi(r1[2])
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	total := 0
	for i := 0; i < n-1; i++ {
		a, _ := strconv.Atoi(r2[i])
		total += a
	}
	x := n*m - total
	if x > k {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, Max(0, x))
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
