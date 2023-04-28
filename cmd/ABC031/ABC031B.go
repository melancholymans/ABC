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
	l, _ := strconv.Atoi(r[0])
	h, _ := strconv.Atoi(r[1])

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		if a > h {
			fmt.Fprintln(writer, -1)
		} else if a < l {
			fmt.Fprintln(writer, l-a)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
}
