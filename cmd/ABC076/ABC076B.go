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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())
	x := 1
	for i := 0; i < n; i++ {
		x = Min(x*2, x+k)
	}
	fmt.Fprintln(writer, x)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
