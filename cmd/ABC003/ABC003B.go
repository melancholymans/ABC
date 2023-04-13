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
	sc.Scan()
	t := sc.Text()
	for i := 0; i < len(s); i++ {
		if s[i] == '@' {
			if strings.Contains("atcoder", string(t[i])) {
				continue
			}
		}
		if t[i] == '@' {
			if strings.Contains("atcoder", string(s[i])) {
				continue
			}
		}
		if s[i] != t[i] {
			fmt.Fprintln(writer, "You will lose")
			return
		}
	}
	fmt.Fprintln(writer, "You can win")
}
