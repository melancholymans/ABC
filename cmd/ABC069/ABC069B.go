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
	fmt.Fprintln(writer, string(s[0])+strconv.Itoa(len(s)-2)+string(s[len(s)-1]))
}
