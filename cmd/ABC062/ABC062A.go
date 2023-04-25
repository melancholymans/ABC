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
	y, _ := strconv.Atoi(r[1])
	a := []int{1, 3, 5, 7, 8, 10, 12}
	b := []int{4, 6, 9, 11}
	c := []int{2}
	if find(a, x) && find(a, y) {
		fmt.Fprintln(writer, "Yes")
		return
	}
	if find(b, x) && find(b, y) {
		fmt.Fprintln(writer, "Yes")
		return
	}
	if find(c, x) && find(c, y) {
		fmt.Fprintln(writer, "Yes")
		return
	}
	fmt.Fprintln(writer, "No")
}

func find(sl []int, t int) bool {
	for i := 0; i < len(sl); i++ {
		if sl[i] == t {
			return true
		}
	}
	return false
}
