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
	h, _ := strconv.Atoi(r1[0])
	w, _ := strconv.Atoi(r1[1])
	sl := make([][]string, 0)
	for i, j := 0, 0; i < h; i++ {
		sc.Scan()
		r2 := strings.Split(sc.Text(), "")
		if calc(r2) {
			sl[j] = make([]string, r2)
			j += 1
		}
	}
	fmt.Fprintln(writer, sl, "\n", w)
}

func calc(r []string) bool {
	for i := 0; i < len(r); i++ {
		if r[i] == "#" {
			return true
		}
	}
	return false
}

func Sli(l, m int, def int) [][]int {
	sl := make([][]int, l)
	for i := 0; i < l; i++ {
		sl[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sl[i][j] = def
		}
	}
	return sl
}
