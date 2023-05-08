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
	total := 700
	for _, v := range s {
		if v == 'o' {
			total += 100
		}
	}
	fmt.Fprintln(writer, total)
}
