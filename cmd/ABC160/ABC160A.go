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
	if s[2] == s[3] && s[4] == s[5] {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
