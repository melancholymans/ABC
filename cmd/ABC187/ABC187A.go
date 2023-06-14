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
	totala := 0
	totalb := 0
	for i := 0; i < 3; i++ {
		totala += a % 10
		totalb += b % 10
		a = a / 10
		b = b / 10
	}
	fmt.Fprintln(writer, Max(totala, totalb))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
