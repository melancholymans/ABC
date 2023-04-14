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
	n, _ := strconv.Atoi(r[0])
	x, _ := strconv.Atoi(r[1])
	a := x - 1
	if n-x < a {
		a = n - x
	}
	fmt.Fprintln(writer, a)
}
