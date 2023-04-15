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
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == t[len(t)-1] {
			for j := len(t) - 1; j >= 0; j-- {
				if string(s[i]) != string(t[j]) {
					count += 1
				}
				i -= 1
			}
			fmt.Fprintln(writer, count)
			return
		}
	}
}
