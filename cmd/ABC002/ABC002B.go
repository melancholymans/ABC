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
	a := make([]byte, 0)
	var b byte
	for i := 0; i < len(s); i++ {
		b = s[i]
		if string(b) != "a" && string(b) != "i" && string(b) != "u" && string(b) != "e" && string(b) != "o" {
			a = append(a, b)
		}
	}
	fmt.Fprintln(writer, string(a))
}
