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
	x, _ := strconv.Atoi(sc.Text())
	m := 100
	for i := 1; ; i++ {
		m = m + (m / 100)
		if x <= m {
			fmt.Fprintln(writer, i)
			break
		}
	}
}
