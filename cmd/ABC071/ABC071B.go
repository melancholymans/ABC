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
	m := [26]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]-97] += 1
	}
	for i := 0; i < 26; i++ {
		if m[i] == 0 {
			fmt.Fprintln(writer, string(i+97))
			return
		}
	}
	fmt.Fprintln(writer, "None")
}
