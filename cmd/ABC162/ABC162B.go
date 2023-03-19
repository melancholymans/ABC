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
	total := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			continue
		}
		if i%5 == 0 {
			continue
		}
		total += i
	}
	fmt.Fprintln(writer, total)
}
