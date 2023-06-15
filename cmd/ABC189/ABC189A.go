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
	c := sc.Text()
	if c[0] == c[1] && c[1] == c[2] {
		fmt.Fprintln(writer, "Won")
	} else {
		fmt.Fprintln(writer, "Lost")
	}
}
