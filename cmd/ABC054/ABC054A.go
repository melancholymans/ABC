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
	a, _ := strconv.Atoi(r[0])
	b, _ := strconv.Atoi(r[1])
	if a == 1 {
		a = 14
	}
	if b == 1 {
		b = 14
	}
	if a > b {
		fmt.Fprintln(writer, "Alice")
	} else if b > a {
		fmt.Fprintln(writer, "Bob")
	} else {
		fmt.Fprintln(writer, "Draw")
	}
}
