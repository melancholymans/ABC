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
	x := sc.Text()
	m := make(map[string]int)
	m["A"] = 1
	m["B"] = 2
	m["C"] = 3
	m["D"] = 4
	m["E"] = 5
	fmt.Fprintln(writer, m[x])
}
