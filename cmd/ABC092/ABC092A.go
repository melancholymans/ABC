package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	c, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	d, _ := strconv.Atoi(sc.Text())
	fmt.Fprintln(writer, Min(a, b)+Min(c, d))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
