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
	t, _ := strconv.Atoi(r[1])
	a := x - t
	if a < 0 {
		a = 0
	}
	fmt.Fprintln(writer, a)
}
