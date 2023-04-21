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
	if x < 1200 {
		fmt.Fprintln(writer, "ABC")
	} else {
		fmt.Fprintln(writer, "ARC")
	}
}
