package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	x, _ := strconv.Atoi(sc.Text())
	a := x / 500
	x = x % 500
	b := x / 5
	fmt.Fprintln(writer, a*1000+b*5)
}
