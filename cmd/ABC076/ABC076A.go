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
	r, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	g, _ := strconv.Atoi(sc.Text())
	fmt.Fprintln(writer, r+(g-r)*2)
}
