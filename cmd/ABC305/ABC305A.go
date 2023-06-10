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
	if n == 100 {
		fmt.Fprintln(writer, 0)
		return
	}
	if n == 0 {
		fmt.Fprintln(writer, 5)
		return
	}
	if n%5 == 0 {
		fmt.Fprintln(writer, n/5*5)
	} else {
		fmt.Fprintln(writer, n/5*5+5)
	}
}
