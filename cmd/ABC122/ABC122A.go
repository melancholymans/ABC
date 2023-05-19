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
	b := sc.Text()
	m := map[string]string{"A": "T", "T": "A", "C": "G", "G": "C"}
	fmt.Fprintln(writer, m[b])
}
