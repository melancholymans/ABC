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
	if s[0] == s[1] && s[0] == s[2] && s[0] == s[3] {
		fmt.Fprintln(writer, "SAME")
	} else {
		fmt.Fprintln(writer, "DIFFERENT")
	}
}
