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
	switch n, _ := strconv.Atoi(r); {
	case 100 == n:
		fmt.Fprintln(writer, "Perfect")
	case 90 <= n && 100 > n:
		fmt.Fprintln(writer, "Great")
	case 60 <= n && 90 > n:
		fmt.Fprintln(writer, "Good")
	default:
		fmt.Fprintln(writer, "Bad")
	}
}
