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
	if x < a {
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintln(writer, 10)
	}
}
