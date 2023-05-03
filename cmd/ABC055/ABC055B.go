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
	p := 1
	for i := 1; i <= n; i++ {
		p = (p * i) % 1000000007
	}
	fmt.Fprintln(writer, p)
}
