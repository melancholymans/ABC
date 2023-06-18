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
	a, _ := strconv.ParseFloat(r[0], 64)
	b, _ := strconv.ParseFloat(r[1], 64)
	fmt.Fprintln(writer, (a-b)*100/a)
}
