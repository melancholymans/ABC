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
	//n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := sc.Text()
	fmt.Fprintln(writer, strings.Count(s, "ABC"))
}
