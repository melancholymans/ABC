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
	switch c[0] {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Fprintln(writer, "vowel")
	default:
		fmt.Fprintln(writer, "consonant")
	}
}
