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
	a, _ := strconv.ParseFloat(r[0], 64)
	b, _ := strconv.ParseFloat(r[1], 64)
	c, _ := strconv.ParseFloat(r[2], 64)
	d, _ := strconv.ParseFloat(r[3], 64)
	if b/a > d/c {
		fmt.Fprintln(writer, "TAKAHASHI")
	} else if b/a < d/c {
		fmt.Fprintln(writer, "AOKI")
	} else {
		fmt.Fprintln(writer, "DRAW")
	}
}
