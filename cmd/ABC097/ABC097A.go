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
	d, _ := strconv.Atoi(r[3])
	if Abs(a-c) <= d || ((Abs(a-b)) <= d && Abs(b-c) <= d) {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
