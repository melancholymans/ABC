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
	c1 := sc.Text()
	sc.Scan()
	c2 := sc.Text()
	if c1[0] == c2[2] && c1[2] == c2[0] && c1[1] == c2[1] {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}

}
