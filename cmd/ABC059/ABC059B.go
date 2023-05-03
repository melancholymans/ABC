package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	a := sc.Text()
	sc.Scan()
	b := sc.Text()
	fa := fmt.Sprintf("%0110s", a)
	fb := fmt.Sprintf("%0110s", b)
	rst := strings.Compare(fa, fb)
	if rst == 1 {
		fmt.Fprintln(writer, "GREATER")
	} else if rst == -1 {
		fmt.Fprintln(writer, "LESS")
	} else if rst == 0 {
		fmt.Fprintln(writer, "EQUAL")
	}
}
