package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r := sc.Text()
	if len(r) == 1 {
		if r[0] == 'a' {
			fmt.Fprintln(writer, -1)
			return
		} else {
			fmt.Fprintln(writer, string(r[0]-1))
			return
		}
	} else {
		fmt.Fprintln(writer, r[0:len(r)-1])
		return
	}
}
