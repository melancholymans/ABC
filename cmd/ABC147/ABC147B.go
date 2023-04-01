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
	s := []byte(sc.Text())
	length := len(s)
	count := 0
	for i, j := 0, len(s)-1; i < length/2; i, j = i+1, j-1 {
		if s[i] != s[j] {
			s[i] = s[j]
			count += 1
		}
	}
	fmt.Fprintln(writer, count)
}
