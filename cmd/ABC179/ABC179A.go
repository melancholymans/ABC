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
	l := len(s)
	if string(s[l-1]) == "s" {
		fmt.Fprintln(writer, s+"es")
	} else {
		fmt.Fprintln(writer, s+"s")
	}
}
