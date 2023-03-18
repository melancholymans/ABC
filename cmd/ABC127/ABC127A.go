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
	s := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[1])
	calc(writer, a, b)
}

func calc(writer *bufio.Writer, a, b int) {
	if a <= 5 {
		fmt.Fprintln(writer, 0)
		return
	}
	if a >= 6 && a <= 12 {
		fmt.Fprintln(writer, b/2)
		return
	}
	fmt.Fprintln(writer, b)
}
