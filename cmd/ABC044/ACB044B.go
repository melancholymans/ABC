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
	w := sc.Text()
	m := map[rune]int{}
	for _, v := range w {
		m[v] += 1
	}
	for _, v := range m {
		if v%2 != 0 {
			fmt.Fprintln(writer, "No")
			return
		}
	}
	fmt.Fprintln(writer, "Yes")
}
