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
	a, _ := strconv.Atoi(r1[0])
	b, _ := strconv.Atoi(r1[1])
	c, _ := strconv.Atoi(r1[2])
	k, _ := strconv.Atoi(r1[3])
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	s, _ := strconv.Atoi(r2[0])
	t, _ := strconv.Atoi(r2[1])
	if k <= s+t {
		a = a - c
		b = b - c
	}
	fmt.Fprintln(writer, s*a+t*b)
}
