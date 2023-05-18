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
	r := strings.Split(sc.Text(), " ")
	t, _ := strconv.ParseFloat(r[0], 64)
	x, _ := strconv.ParseFloat(r[1], 64)
	fmt.Fprintln(writer, t/x)
}
