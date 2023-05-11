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
	total := 0
	for _, v := range s {
		if v == '+' {
			total += 1
		} else {
			total -= 1
		}
	}
	fmt.Fprintln(writer, total)
}
