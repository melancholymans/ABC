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
	k, _ := strconv.Atoi(r[2])
	if a < k {
		k = k - a
		a = 0
	} else {
		a = a - k
		fmt.Fprintln(writer, a, b)
		return
	}
	if b < k {
		b = 0
		fmt.Fprintln(writer, a, b)
		return
	} else {
		b = b - k
		fmt.Fprintln(writer, a, b)
	}
}
