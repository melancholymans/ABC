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
	a, _ := strconv.Atoi(sc.Text())
	fmt.Fprintln(writer, a*(1+a+a*a))
}
