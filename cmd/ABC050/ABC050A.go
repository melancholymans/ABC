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
	b, _ := strconv.Atoi(r[2])
	op := r[1]
	if op == "+" {
		fmt.Fprintln(writer, a+b)
	} else {
		fmt.Fprintln(writer, a-b)
	}
}
