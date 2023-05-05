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
	r1 := sc.Text()
	n, _ := strconv.Atoi(r1)
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	ad := make([]int, 0)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(r2[i])
		ad = append(ad, a)
	}
	fmt.Fprintln(writer, n)
	fmt.Fprintln(writer, ad)
}
