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
	s := sc.Text()
	if s == "Hello,World!" {
		fmt.Fprintln(writer, "AC")
	} else {
		fmt.Fprintln(writer, "WA")
	}
}
