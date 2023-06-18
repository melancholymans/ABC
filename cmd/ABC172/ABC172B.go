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
	s := sc.Text()
	sc.Scan()
	t := sc.Text()
	n := len(s)
	count := 0
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			count += 1
		}
	}
	fmt.Fprintln(writer, count)
}
