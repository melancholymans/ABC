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
	s := sc.Text()
	sl := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			sl = append(sl, string(s[i]))
		}
	}
	fmt.Fprintln(writer, strings.Join(sl, ""))
}
