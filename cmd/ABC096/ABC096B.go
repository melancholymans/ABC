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
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())
	d := Max(Max(a, b), c)
	e := (a + b + c) - d
	for i := 0; i < k; i++ {
		d = d * 2
	}
	fmt.Fprintln(writer, d+e)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
