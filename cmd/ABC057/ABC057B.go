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
	for i := 0; i < n; i++ {
		sc.Scan()
		r2 := strings.Split(sc.Text(), "")
		fmt.Fprintln(writer, r2)
	}
	for j := 0; j < m; j++ {
		sc.Scan()
		r3 := strings.Split(sc.Text(), "")
		fmt.Fprintln(writer, r3)
	}
	fmt.Fprintln(writer, n, m)
}
