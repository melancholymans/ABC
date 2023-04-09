package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	count := 0
	for i := 0; i < 12; i++ {
		sc.Scan()
		s := sc.Text()
		if strings.Contains(s, "r") {
			count += 1
		}
	}
	fmt.Fprintln(writer, count)
}
