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
	a := r[0]
	b := r[1]
	if a == "H" && b == "H" {
		fmt.Fprintln(writer, "H")
	}
	if a == "H" && b == "D" {
		fmt.Fprintln(writer, "D")
	}
	if a == "D" && b == "H" {
		fmt.Fprintln(writer, "D")
	}
	if a == "D" && b == "D" {
		fmt.Fprintln(writer, "H")
	}
}
