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
	total := 0.0
	for i := 0; i < n; i++ {
		sc.Scan()
		r2 := strings.Split(sc.Text(), " ")
		x, _ := strconv.ParseFloat(r2[0], 64)
		if r2[1] == "JPY" {
			total += x
		} else {
			total += x * 380000.0
		}
	}
	fmt.Fprintln(writer, total)
}
