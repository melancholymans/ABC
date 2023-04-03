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
	s := sc.Text()
	n, _ := strconv.Atoi(s)
	if n == 1 {
		fmt.Fprintln(writer, "ABC")
	} else {
		fmt.Fprintln(writer, "chokudai")
	}
}
