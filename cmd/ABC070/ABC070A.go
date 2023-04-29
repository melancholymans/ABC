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
	rs := []rune(n)
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		if rs[i] != rs[j] {
			fmt.Fprintln(writer, "No")
			return
		}
	}
	fmt.Fprintln(writer, "Yes")
}
