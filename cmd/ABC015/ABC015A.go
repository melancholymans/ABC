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
	a := sc.Text()
	sc.Scan()
	b := sc.Text()
	if len(a) > len(b) {
		fmt.Fprintln(writer, a)
	} else {
		fmt.Fprintln(writer, b)
	}
}
