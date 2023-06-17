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
	if n%2 == 0 {
		fmt.Fprintln(writer, "White")
	} else {
		fmt.Fprintln(writer, "Black")
	}
}
