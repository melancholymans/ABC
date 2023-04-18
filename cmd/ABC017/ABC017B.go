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
	sc.Scan()
	x := sc.Text()
	x = strings.ReplaceAll(x, "ch", "")
	x = strings.ReplaceAll(x, "o", "")
	x = strings.ReplaceAll(x, "k", "")
	x = strings.ReplaceAll(x, "u", "")
	if len(x) > 0 {
		fmt.Fprintln(writer, "NO")
	} else {
		fmt.Fprintln(writer, "YES")
	}
}
