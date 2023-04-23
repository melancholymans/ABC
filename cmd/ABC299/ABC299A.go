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
	s := sc.Text()
	s = strings.ReplaceAll(s, ".", "")
	n = n / 10
	if s == "|*|" {
		fmt.Fprintln(writer, "in")
	} else {
		fmt.Fprintln(writer, "out")
	}
}
