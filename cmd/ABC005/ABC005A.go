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
	s := strings.Split(sc.Text(), " ")
	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])
	fmt.Fprintln(writer, y/x)
}
