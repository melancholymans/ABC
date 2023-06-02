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
	if a == b && b != c {
		fmt.Fprintln(writer, "Yes")
		return
	}
	if b == c && c != a {
		fmt.Fprintln(writer, "Yes")
		return
	}
	if c == a && a != b {
		fmt.Fprintln(writer, "Yes")
		return
	}
	fmt.Fprintln(writer, "No")
}
