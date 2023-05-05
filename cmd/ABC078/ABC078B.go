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
	y, _ := strconv.Atoi(r[1])
	z, _ := strconv.Atoi(r[2])
	d := x / (y + z)
	if d*(y+z)+z == x {
		fmt.Fprintln(writer, d)
	} else {
		fmt.Fprintln(writer, d-1)
	}
}
