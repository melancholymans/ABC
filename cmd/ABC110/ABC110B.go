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
	n, _ := strconv.Atoi(r1[0])
	m, _ := strconv.Atoi(r1[1])
	x, _ := strconv.Atoi(r1[2])
	y, _ := strconv.Atoi(r1[3])
	xmax := x
	ymin := y
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(r2[i])
		if xmax < a {
			xmax = a
		}
	}
	sc.Scan()
	r3 := strings.Split(sc.Text(), " ")
	for i := 0; i < m; i++ {
		a, _ := strconv.Atoi(r3[i])
		if ymin > a {
			ymin = a
		}
	}
	if xmax < ymin {
		fmt.Fprintln(writer, "No War")
	} else {
		fmt.Fprintln(writer, "War")
	}
}
