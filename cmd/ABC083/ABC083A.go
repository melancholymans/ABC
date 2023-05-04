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
	if a+b > c+d {
		fmt.Fprintln(writer, "Left")
	} else if a+b < c+d {
		fmt.Fprintln(writer, "Right")
	} else {
		fmt.Fprintln(writer, "Balanced")
	}
}
