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
	l1, _ := strconv.Atoi(r[0])
	l2, _ := strconv.Atoi(r[1])
	l3, _ := strconv.Atoi(r[2])
	result := 0
	if l1 == l2 {
		result = l3
	}
	if l1 == l3 {
		result = l2
	}
	if l2 == l3 {
		result = l1
	}
	fmt.Fprintln(writer, result)
}
