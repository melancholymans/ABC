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
	total := 0
	for i := 0; i < n; i++ {
		total += i + 1
	}
	fmt.Fprintln(writer, total*10000/n)
}
