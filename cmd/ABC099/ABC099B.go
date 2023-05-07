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
	fmt.Fprintln(writer, calc(b-a)-b)
}

func calc(a int) int {
	total := 0
	for i := 0; i < a; i++ {
		total += i + 1
	}
	return total
}
