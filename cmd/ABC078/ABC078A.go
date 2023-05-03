package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r := strings.Split(sc.Text(), " ")
	x := r[0]
	y := r[1]
	if x > y {
		fmt.Fprintln(writer, ">")
	} else if x < y {
		fmt.Fprintln(writer, "<")
	} else {
		fmt.Fprintln(writer, "=")
	}
}
