package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	o := sc.Text()
	sc.Scan()
	e := sc.Text()
	a := make([]byte, 0)
	lim := Max(len(o), len(e))
	for i := 0; i < lim; i++ {
		if i < len(o) {
			a = append(a, o[i])
		}
		if i < len(e) {
			a = append(a, e[i])
		}
	}
	fmt.Fprintln(writer, string(a))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
