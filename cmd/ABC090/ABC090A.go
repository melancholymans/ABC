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
	sc.Scan()
	c3 := sc.Text()
	fmt.Fprintln(writer, string(c1[0])+string(c2[1])+string(c3[2]))
}
