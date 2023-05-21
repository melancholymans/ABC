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
	m, _ := strconv.Atoi(r1[1])
	c, _ := strconv.Atoi(r1[2])
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	slb := make([]int, 0)
	for i := 0; i < m; i++ {
		b, _ := strconv.Atoi(r2[i])
		slb = append(slb, b)
	}
	count := 0
	for i := 0; i < n; i++ {
		total := 0
		sc.Scan()
		r3 := strings.Split(sc.Text(), " ")
		for j := 0; j < m; j++ {
			a, _ := strconv.Atoi(r3[j])
			total += slb[j] * a
		}
		if total+c > 0 {
			count += 1
		}
	}
	fmt.Fprintln(writer, count)
}
