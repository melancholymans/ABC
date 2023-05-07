package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r := strings.Split(sc.Text(), " ")
	d, _ := strconv.Atoi(r[0])
	n, _ := strconv.Atoi(r[1])
	if d == 0 {
		fmt.Fprintln(writer, n+n/100)
	} else if d == 1 {
		fmt.Fprintln(writer, 100*n)
	} else {
		fmt.Fprintln(writer, 10000*n)
	}
}
