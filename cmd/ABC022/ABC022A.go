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
	n, _ := strconv.Atoi(r[0])
	s, _ := strconv.Atoi(r[1])
	t, _ := strconv.Atoi(r[2])
	sc.Scan()
	w, _ := strconv.Atoi(sc.Text())
	count := 0
	if w >= s && w <= t {
		count += 1
	}
	for i := 0; i < n-1; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		w += a
		if w >= s && w <= t {
			count += 1
		}
	}
	fmt.Fprintln(writer, count)
}
