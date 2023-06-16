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
	m, _ := strconv.ParseFloat(r1[1], 64)
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	total := 0.0
	sl := make([]float64, n)
	for i := 0; i < n; i++ {
		a, _ := strconv.ParseFloat(r2[i], 64)
		sl[i] = a
		total += a
	}
	count := 0
	for i := 0; i < n; i++ {
		if 1.0/(4*m) <= sl[i]/total {
			count += 1
		}
	}
	if count >= int(m) {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
