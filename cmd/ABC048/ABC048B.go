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
	a, _ := strconv.Atoi(r[0])
	b, _ := strconv.Atoi(r[1])
	x, _ := strconv.Atoi(r[2])
	if b < x {
		fmt.Fprintln(writer, 1)
	} else {
		fmt.Fprintln(writer, b/x+1-((a-1)/x+1))
	}
}
