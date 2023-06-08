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
	a := sc.Text()[0]
	if 'A' <= a && a <= 'Z' {
		fmt.Fprintln(writer, "A")
		return
	}
	if 'a' <= a && a <= 'z' {
		fmt.Fprintln(writer, "a")
		return
	}
}
