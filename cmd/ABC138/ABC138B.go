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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	r := strings.Split(sc.Text(), " ")
	total := 0.0
	for i := 0; i < n; i++ {
		a, _ := strconv.ParseFloat(r[i], 64)
		total += 1 / a
	}
	fmt.Fprintln(writer, 1/total)
}
