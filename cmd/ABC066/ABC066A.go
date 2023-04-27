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
	r := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(r[0])
	b, _ := strconv.Atoi(r[1])
	c, _ := strconv.Atoi(r[2])
	x := a + b
	y := b + c
	z := c + a
	fmt.Fprintln(writer, Min(Min(x, y), z))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
