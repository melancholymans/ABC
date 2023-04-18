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
	c, _ := strconv.Atoi(r[2])
	if a+b == c && a-b == c {
		fmt.Fprintln(writer, "?")
	} else if a+b == c {
		fmt.Fprintln(writer, "+")
	} else if a-b == c {
		fmt.Fprintln(writer, "-")
	} else {
		fmt.Fprintln(writer, "!")
	}
}
