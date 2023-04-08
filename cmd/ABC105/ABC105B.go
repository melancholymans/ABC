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
	n, _ := strconv.Atoi(r)
	for k := 0; k < n/4+1; k++ {
		for d := 0; d < n/7+1; d++ {
			if k*4+d*7 == n {
				fmt.Fprintln(writer, "Yes")
				return
			}
		}
	}
	fmt.Fprintln(writer, "No")
}
