package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	n := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(n[0])
	b, _ := strconv.Atoi(n[1])
	sc.Scan()
	s := sc.Text()
	if !calc(a, b, s) {
		fmt.Fprintln(writer, "No")
	} else {
		fmt.Fprintln(writer, "Yes")
	}
}

func calc(a, b int, s string) bool {
	if len(s) != a+b+1 {
		return false
	}
	if s[a] != '-' {
		return false
	}
	if !numCheck(s[:a]) {
		return false
	}
	if !numCheck(s[a+1:]) {
		return false
	}
	return true
}

func numCheck(s string) bool {
	for _, r := range s {
		if '0' > r || r > '9' {
			return false
		}
	}
	return true
}
