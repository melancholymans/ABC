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
	switch n % 10 {
	case 2, 4, 5, 7, 9:
		fmt.Fprintln(writer, "hon")
	case 0, 1, 6, 8:
		fmt.Fprintln(writer, "pon")
	case 3:
		fmt.Fprintln(writer, "bon")
	}
}
