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
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := sc.Text()
	if a >= 3200 {
		fmt.Fprintln(writer, s)
	} else {
		fmt.Fprintln(writer, "red")
	}
}
