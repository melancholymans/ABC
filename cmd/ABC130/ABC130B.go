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
	s := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(s[0])
	x, _ := strconv.Atoi(s[1])
	sc.Scan()
	la := strings.Split(sc.Text(), " ")
	count := 0
	d := 0
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(la[i])
		d = d + a
		if d > x {
			break
		}
		count += 1
	}
	fmt.Fprintln(writer, count+1)
}
