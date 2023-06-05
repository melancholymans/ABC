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
	n := len(s)
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			if string(s[i]) != "R" {
				continue
			} else {
				fmt.Fprintln(writer, "No")
				return
			}
		} else {
			if string(s[i]) != "L" {
				continue
			} else {
				fmt.Fprintln(writer, "No")
				return
			}
		}
	}
	fmt.Fprintln(writer, "Yes")
}
