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
	w, _ := strconv.Atoi(r1[0])
	h, _ := strconv.Atoi(r1[1])
	n, _ := strconv.Atoi(r1[2])

	for i := 0; i < n; i++ {
		sc.Scan()
		r2 := strings.Split(sc.Text(), " ")
		fmt.Fprintln(writer, r2)
	}
	fmt.Fprintln(writer, w, h, n)
}
