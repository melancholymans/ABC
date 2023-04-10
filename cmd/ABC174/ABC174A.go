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
	r := sc.Text()
	x, _ := strconv.Atoi(r)
	if x >= 30 {
		fmt.Fprintln(writer, "Yes")
		return
	}
	fmt.Fprintln(writer, "No")
}
