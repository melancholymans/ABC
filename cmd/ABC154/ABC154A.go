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
	r1 := strings.Split(sc.Text(), " ")
	s := r1[0]
	t := r1[1]

	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(r2[0])
	b, _ := strconv.Atoi(r2[1])

	sc.Scan()
	u := sc.Text()

	if u == s {
		fmt.Fprintln(writer, a-1, b)
	} else if u == t {
		fmt.Fprintln(writer, a, b-1)
	}
}
