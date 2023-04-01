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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := strings.Split(sc.Text(), "")
	if calc(n, s) {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}

func calc(n int, s []string) bool {
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			return false
		}
	}
	return true
}
