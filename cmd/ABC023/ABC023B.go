package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := sc.Text()
	if n%2 == 0 {
		fmt.Fprintln(writer, -1)
		return
	}
	l := len(s) - 1
	count := 0
	for i := 0; i < n/2; i++ {
		if s[i] == 'b' && s[l-i] == 'b' {
			count += 1
			continue
		} else if s[i] == 'c' && s[l-i] == 'a' {
			count += 1
			continue
		} else if s[i] == 'a' && s[l-i] == 'c' {
			count += 1
			continue
		} else {
			fmt.Fprintln(writer, -1)
			return
		}
	}
	fmt.Fprintln(writer, count)
}
