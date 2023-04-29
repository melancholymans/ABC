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
	x, _ := strconv.Atoi(r[0])
	a, _ := strconv.Atoi(r[1])
	b, _ := strconv.Atoi(r[2])
	if Abs(x-a) < Abs(x-b) {
		fmt.Fprintln(writer, "A")

	} else {
		fmt.Fprintln(writer, "B")
	}
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
