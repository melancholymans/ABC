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
	m := map[rune]int{}
	for _, v := range s {
		m[v] += 1
	}
	fmt.Fprintln(writer, m['A'], m['B'], m['C'], m['D'], m['E'], m['F'])
}
