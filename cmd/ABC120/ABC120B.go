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
	k, _ := strconv.Atoi(r[2])
	n := Max(a, b)
	sl := make([]int, 0)
	for i := 1; i < n+1; i++ {
		if a%i == 0 && b%i == 0 {
			sl = append(sl, i)
		}
	}
	fmt.Fprintln(writer, sl[len(sl)-k])
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
