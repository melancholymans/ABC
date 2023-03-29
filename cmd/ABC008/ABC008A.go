package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	c := strings.Split(sc.Text(), " ")
	s, _ := strconv.Atoi(c[0])
	t, _ := strconv.Atoi(c[1])
	fmt.Fprintln(writer, t-s+1)
}
