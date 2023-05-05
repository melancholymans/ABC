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
	sc.Scan()
	r := strings.Split(sc.Text(), " ")
	m := make(map[string]int, 0)
	for i := 0; i < n; i++ {
		s := r[i]
		m[s] += 1
	}
	if len(m) == 3 {
		fmt.Fprintln(writer, "Three")
	} else if len(m) == 4 {
		fmt.Fprintln(writer, "Four")
	}
}
