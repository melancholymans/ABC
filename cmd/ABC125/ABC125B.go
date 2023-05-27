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
	v := make([]int, 0)
	c := make([]int, 0)
	sc.Scan()
	r1 := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		x, _ := strconv.Atoi(r1[i])
		v = append(v, x)
	}
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		y, _ := strconv.Atoi(r2[i])
		c = append(c, y)
	}
	total := 0
	for i := 0; i < n; i++ {
		if v[i]-c[i] > 0 {
			total += v[i] - c[i]
		}
	}
	fmt.Fprintln(writer, total)
}
