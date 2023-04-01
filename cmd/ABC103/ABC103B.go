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
	sc.Scan()
	t := []byte(sc.Text())
	for i := 0; i < len(s); i++ {
		if string(calc(s)) == string(t) {
			fmt.Fprintln(writer, "Yes")
			goto Z
		}
	}
	fmt.Fprintln(writer, "No")
Z:
}

func calc(b []byte) []byte {
	m := b[len(b)-1]
	for i := len(b) - 1; i > 0; i-- {
		b[i] = b[i-1]
	}
	b[0] = m
	return b
}
