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
	n := sc.Text()
	m := [26]int{}
	for i := 0; i < len(n); i++ {
		m[n[i]-97] += 1
	}
	for i := 0; i < 26; i++ {
		if m[i] > 1 {
			fmt.Fprintln(writer, "no")
			return
		}
	}
	fmt.Fprintln(writer, "yes")
}
