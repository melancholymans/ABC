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
	switch n {
	case 22:
		fmt.Fprintln(writer, "Christmas Eve Eve Eve")
	case 23:
		fmt.Fprintln(writer, "Christmas Eve Eve")
	case 24:
		fmt.Fprintln(writer, "Christmas Eve")
	case 25:
		fmt.Fprintln(writer, "Christmas")
	}
}
