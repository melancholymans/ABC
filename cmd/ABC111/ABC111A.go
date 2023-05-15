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
	n := []byte(sc.Text())
	if n[0] == 57 {
		n[0] = 49
	} else if n[0] == 49 {
		n[0] = 57
	}
	if n[1] == 57 {
		n[1] = 49
	} else if n[1] == 49 {
		n[1] = 57
	}
	if n[2] == 57 {
		n[2] = 49
	} else if n[2] == 49 {
		n[2] = 57
	}
	fmt.Fprintln(writer, string(n))
}
