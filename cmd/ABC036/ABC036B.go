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
	sl := make([][]string, n)
	for i := 0; i < n; i++ {
		sl[i] = make([]string, n)
		sc.Scan()
		r := strings.Split(sc.Text(), "")
		for j := 0; j < n; j++ {
			sl[i][j] = r[j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(writer, sl[n-1-j][i])
		}
		fmt.Fprint(writer, "\n")
	}
}
