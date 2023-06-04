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

	for i := 0; i < 3; i++ {
		sc.Scan()
		r := strings.Split(sc.Text(), " ")
		fmt.Fprintln(writer, r)
	}
	fmt.Fprintln(writer, n)
}
