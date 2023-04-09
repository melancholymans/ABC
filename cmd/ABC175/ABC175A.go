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
	st := sc.Text()
	count := 0
	var m int
	for _, s := range st {
		if string(s) == "R" {
			count += 1
		} else {
			count = 0
		}
		if count > 0 {
			m = count
		}
	}
	fmt.Fprintln(writer, m)
}
