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
	length := len(s)
	for i := 0; i < length; i++ {
		s = s[0 : len(s)-1]
		l := len(s)
		if l%2 == 0 && calc(s[0:l/2], s[l/2:]) {
			fmt.Fprintln(writer, len(s))
			break
		}
	}
}

func calc(s1, s2 string) bool {
	return s1 == s2
}
