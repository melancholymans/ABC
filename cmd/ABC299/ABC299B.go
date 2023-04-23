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
	t, _ := strconv.Atoi(r1[1])
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	cd := make([]int, 0)
	for i := 0; i < n; i++ {
		c, _ := strconv.Atoi(r1[i])
		cd = append(cd, c)
	}
	rd := make([]int, 0)
	for i := 0; i < n; i++ {
		r, _ := strconv.Atoi(r2[i])
		cd = append(rd, r)
	}

	fmt.Fprintln(writer, n, t)
	fmt.Fprintln(writer, cd)
	fmt.Fprintln(writer, rd)
}
