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
	r, _ := strconv.Atoi(sc.Text())

	if 1200 > r {
		fmt.Fprintln(writer, "ABC")
		return
	} else if 2800 > r {
		fmt.Fprintln(writer, "ARC")
		return
	} else {
		fmt.Fprintln(writer, "AGC")
	}
}
