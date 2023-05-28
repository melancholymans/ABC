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
	if calc(a) && calc(b) {
		fmt.Fprintln(writer, a*b)
	} else {
		fmt.Fprintln(writer, -1)
	}
}

func calc(a int) bool {
	if a >= 1 && a <= 9 {
		return true
	}
	return false
}
