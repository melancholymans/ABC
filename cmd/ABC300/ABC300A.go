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
	a, _ := strconv.Atoi(r1[1])
	b, _ := strconv.Atoi(r1[2])
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		c, _ := strconv.Atoi(r2[i])
		if c == a+b {
			fmt.Fprintln(writer, i+1)
			return
		}
	}
}
