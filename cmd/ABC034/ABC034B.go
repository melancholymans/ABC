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
	if n%2 == 0 {
		fmt.Fprintln(writer, n-1)
	} else {
		fmt.Fprintln(writer, n+1)
	}
}
