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
	k, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	r := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(r[i])
		fmt.Fprintln(writer, a)
	}
	fmt.Fprintln(writer, n, k)
}
