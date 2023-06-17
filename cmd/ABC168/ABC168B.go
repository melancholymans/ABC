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
	k, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := sc.Text()
	if k < len(s) {
		fmt.Fprintln(writer, s[:k]+"...")
	} else {
		fmt.Fprintln(writer, s)
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
