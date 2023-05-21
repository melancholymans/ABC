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
	p, _ := strconv.Atoi(r1[0])
	q, _ := strconv.Atoi(r1[1])
	r, _ := strconv.Atoi(r1[2])
	maxm := Max(Max(p, q), r)
	fmt.Fprintln(writer, p+q+r-maxm)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
