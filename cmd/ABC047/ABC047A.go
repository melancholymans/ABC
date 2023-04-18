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
	max := Max(Max(a, b), c)
	total := a + b + c
	if total-max == max {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
