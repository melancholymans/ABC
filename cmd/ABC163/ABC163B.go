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
	n, _ := strconv.Atoi(r[0])
	m, _ := strconv.Atoi(r[1])
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	total := 0
	for i := 0; i < m; i++ {
		a, _ := strconv.Atoi(r2[i])
		total += a
	}
	if total <= n {
		fmt.Fprintln(writer, n-total)
	} else {
		fmt.Fprintln(writer, -1)
	}
}
