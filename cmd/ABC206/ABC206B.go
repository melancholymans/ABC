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
	i := 1
	for {
		total += i
		if total >= n {
			break
		}
		i += 1
	}
	fmt.Fprintln(writer, i)
}
