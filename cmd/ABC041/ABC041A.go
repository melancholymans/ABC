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
	s := sc.Text()
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	fmt.Fprintln(writer, string(s[i-1]))
}
