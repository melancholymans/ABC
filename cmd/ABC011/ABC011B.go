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
	r := sc.Text()
	var s string
	for i, v := range r {
		if v < 97 {
			v = v - 65 + 97
		}
		if i == 0 && v >= 97 {
			v = v - 97 + 65
		}
		s = s + string(v)
	}
	fmt.Fprintln(writer, s)
}
