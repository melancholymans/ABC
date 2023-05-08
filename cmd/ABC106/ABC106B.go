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
	count := 0
	for i := 1; i <= n; i += 2 {
		c := calc(i)
		if c == 8 {
			count += 1
		}
	}
	fmt.Fprintln(writer, count)
}

func calc(n int) int {
	count := 0
	for i := 1; i < n+1; i++ {
		if n%i == 0 {
			count += 1
		}
	}
	return count
}
