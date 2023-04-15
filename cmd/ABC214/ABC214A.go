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
	r := sc.Text()
	n, _ := strconv.Atoi(r)
	if 1 <= n && 125 >= n {
		fmt.Fprintln(writer, 4)
	} else if 126 <= n && 211 >= n {
		fmt.Fprintln(writer, 6)
	} else {
		fmt.Fprintln(writer, 8)
	}
}
